package service

import (
	"sync"

	"github.com/mental-health/model"
	"github.com/mental-health/util"

	"github.com/lexkong/log"
)

type ParentCommentInfoList struct {
	Lock  *sync.Mutex
	IdMap map[uint32]*model.ParentCommentInfo2
}

type SubCommentInfoList struct {
	Lock  *sync.Mutex
	IdMap map[uint32]*model.CommentInfo
}

// Get comment list.
func CommentList(holeId uint32, limit, offset int32, userId uint32) (*[]model.ParentCommentInfo2, error) {
	// Get parent comments from database
	parentComments, err := model.GetParentComments(holeId, limit, offset)
	if err != nil {
		log.Error("GetParentComments", err)
		return nil, err
	}

	var parentIds []uint32
	for _, parentComment := range *parentComments {
		parentIds = append(parentIds, parentComment.Id)
	}

	parentCommentInfoList := ParentCommentInfoList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint32]*model.ParentCommentInfo2, len(*parentComments)),
	}

	wg := new(sync.WaitGroup)
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, parentComment := range *parentComments {
		wg.Add(1)
		go func(parentComment model.ParentCommentModel) {
			defer wg.Done()

			// 获取父评论详情
			parentCommentInfo2, err := GetParentCommentInfo2(parentComment.Id, userId)
			if err != nil {
				log.Error("GetParentCommentInfo function error", err)
				errChan <- err
				return
			}

			parentCommentInfoList.Lock.Lock()
			defer parentCommentInfoList.Lock.Unlock()

			parentCommentInfoList.IdMap[parentCommentInfo2.Id] = parentCommentInfo2

		}(parentComment)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}

	var infos []model.ParentCommentInfo2
	for _, id := range parentIds {
		infos = append(infos, *parentCommentInfoList.IdMap[id])
	}

	return &infos, nil
}

// Get the response data information of a parent comment.
func GetParentCommentInfo2(id uint32, userId uint32) (*model.ParentCommentInfo2, error) {
	// Get comment from Database
	comment := &model.ParentCommentModel{Id: id}
	if err := comment.GetById(); err != nil {
		log.Error("comment.GetById() error", err)
		return nil, err
	}

	user := &model.UserModel{Id: comment.UserId}
	if err := user.GetUserById(); err != nil {
		log.Error("user.GetUserById() error", err)
		return nil, err
	}

	// Get the user of the parent comment
	userInfo := &model.UserHoleResponse2{
		Username:  user.Username,
		Sid:       user.Sid,
		Avatar:    user.Avatar,
		IsTeacher: user.IsTeacher,
	}
	var err error

	// Get like state
	_, isLike := model.CommentHasLiked(userId, comment.Id)

	// Get subComments' infos
	subCommentInfos, err := GetSubCommentInfosByParentId(comment.Id, userId)
	if err != nil {
		log.Error("GetSubCommentInfosByParentId() error", err)
		return nil, err
	}

	likeNum, err := model.GetCommentLikeSum(comment.Id)
	if err != nil {
		log.Error("model.GetCommentLikeSum() error", err)
		return nil, err
	}

	data := &model.ParentCommentInfo2{
		Id:              comment.Id,
		Content:         comment.Content,
		LikeNum:         likeNum,
		IsLike:          isLike,
		Time:            util.GetCurrentTime(),
		UserInfo:        userInfo,
		SubCommentsNum:  comment.SubCommentNum,
		SubCommentsList: subCommentInfos,
	}

	return data, nil
}

// Get subComments' infos by parent id.
func GetSubCommentInfosByParentId(id uint32, userId uint32) (*[]model.CommentInfo, error) {
	// Get subComments from Database
	comments, err := model.GetSubCommentsByParentId(id)
	if err != nil {
		log.Error("GetSubCommentsByParentId function error", err)
		return nil, err
	}

	var commentIds []uint32
	for _, comment := range *comments {
		commentIds = append(commentIds, comment.Id)
	}

	subCommentInfoList := SubCommentInfoList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint32]*model.CommentInfo, len(*comments)),
	}

	wg := new(sync.WaitGroup)
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// 并发获取子评论详情列表
	for _, comment := range *comments {
		wg.Add(1)

		go func(comment model.SubCommentModel) {
			defer wg.Done()

			// Get a subComment's info by its id
			info, err := GetSubCommentInfoById(comment.Id, userId)
			if err != nil {
				log.Error("GetSubCommentInfoById function error", err)
				errChan <- err
				return
			}

			subCommentInfoList.Lock.Lock()
			defer subCommentInfoList.Lock.Unlock()

			subCommentInfoList.IdMap[info.Id] = info

		}(comment) //传址会panic
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}

	var commentInfos []model.CommentInfo
	for _, id := range commentIds {
		commentInfos = append(commentInfos, *subCommentInfoList.IdMap[id])
	}

	return &commentInfos, nil
}

// Get the response information of a subComment by id.
func GetSubCommentInfoById(id uint32, userId uint32) (*model.CommentInfo, error) {
	// Get comment from Database
	comment := &model.SubCommentModel{Id: id}
	if err := comment.GetById(); err != nil {
		log.Error("comment.GetById function error", err)
		return nil, err
	}

	// Get the user of the subComment if not anonymous
	var commentUser *model.UserInfoResponse
	var err error
	commentUser, err = model.GetUserInfoById(comment.UserId)
	if err != nil {
		log.Error("GetUserInfoById function error", err)
		return nil, err
	}

	// Get the target user of the subComment if not anonymous (identified by 0)
	var targetUser *model.UserInfoResponse
	if comment.TargetUserId != 0 {
		targetUser, err = model.GetUserInfoById(comment.TargetUserId)
		if err != nil {
			log.Error("GetUserInfoById function error", err)
			return nil, err
		}
	}

	// Get like state
	_, isLike := model.CommentHasLiked(userId, comment.Id)

	likeNum, err := model.GetCommentLikeSum(comment.Id)
	if err != nil {
		log.Error("model.GetCommentLikeSum() error", err)
		return nil, err
	}

	data := &model.CommentInfo{
		Id:             comment.Id,
		Content:        comment.Content,
		LikeNum:        likeNum,
		IsLike:         isLike,
		Time:           util.GetCurrentTime(),
		UserInfo:       commentUser,
		TargetUserInfo: targetUser,
	}

	return data, nil
}
