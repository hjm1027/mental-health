package service

import (
	"github.com/lexkong/log"

	"github.com/mental-health/model"
)

func GetCollectionList(userId, limit, page uint32) ([]*model.HoleInfoResponse, error) {
	var response []*model.HoleInfoResponse

	records, err := model.GetHoleCollectionsByUserId(userId, limit, page)
	if err != nil {
		log.Error("GetHoleCollectionsByUserId get records error", err)
		return nil, err
	}

	var holeIds []uint32
	for _, record := range *records {
		holeIds = append(holeIds, record.HoleId)
	}
	//fmt.Println(holeIds)

	for i := 0; i < len(holeIds); i++ {
		//fmt.Println(holeIds[i])
		//get hole
		hole := model.HoleModel{Id: holeIds[i]}
		//fmt.Println(hole)
		if err = hole.GetById(); err != nil {
			log.Error("GetCollectionsByUserId getbyid error", err)
			return nil, err
		}

		//get user
		user, err := model.GetUserInfoById(userId)
		if err != nil {
			log.Error("GetCollectionsByUserId get user error", err)
		}
		userInfo := model.UserHoleResponse{
			Username: user.Username,
			Avatar:   user.Avatar,
			Sid:      user.Sid,
		}

		//get like state
		_, isLike := hole.HasLiked(userId)
		_, isFavorite := hole.HasFavorited(userId)

		data := &model.HoleInfoResponse{
			HoleId:      hole.Id,
			Type:        hole.Type,
			Content:     hole.Content,
			LikeNum:     hole.LikeNum,
			ReadNum:     hole.ReadNum,
			FavoriteNum: hole.FavoriteNum,
			IsLike:      isLike,
			IsFavorite:  isFavorite,
			Time:        hole.Time,
			CommentNum:  hole.CommentNum,
			UserInfo:    userInfo,
		}

		response = append(response, data)
	}
	return response, nil
}

func GetHoleList(userId, limit, page uint32) ([]*model.HoleInfoResponse, error) {
	var response []*model.HoleInfoResponse

	records, err := model.GetHoleList(userId, limit, page)
	if err != nil {
		log.Error("GetHoleList get records error", err)
		return nil, err
	}

	var holeIds []uint32
	for _, record := range *records {
		holeIds = append(holeIds, record.Id)
	}
	//fmt.Println(holeIds)

	for i := 0; i < len(holeIds); i++ {
		//fmt.Println(holeIds[i])
		//get hole
		hole := model.HoleModel{Id: holeIds[i]}
		//fmt.Println(hole)
		if err = hole.GetById(); err != nil {
			log.Error("GetHoleList getbyid error", err)
			return nil, err
		}

		//get user
		user, err := model.GetUserInfoById(userId)
		if err != nil {
			log.Error("GetHoleList get user error", err)
		}
		userInfo := model.UserHoleResponse{
			Username: user.Username,
			Avatar:   user.Avatar,
			Sid:      user.Sid,
		}

		//get like state
		_, isLike := hole.HasLiked(userId)
		_, isFavorite := hole.HasFavorited(userId)

		data := &model.HoleInfoResponse{
			HoleId:      hole.Id,
			Type:        hole.Type,
			Content:     hole.Content,
			LikeNum:     hole.LikeNum,
			ReadNum:     hole.ReadNum,
			FavoriteNum: hole.FavoriteNum,
			IsLike:      isLike,
			IsFavorite:  isFavorite,
			Time:        hole.Time,
			CommentNum:  hole.CommentNum,
			UserInfo:    userInfo,
		}

		response = append(response, data)
	}
	return response, nil
}

// 新建父评论时获取信息
func GetParentCommentInfo(id uint32, userId uint32) (*model.ParentCommentInfo, error) {
	// Get comment from Database
	comment := &model.ParentCommentModel{Id: id}
	if err := comment.GetById(); err != nil {
		log.Error("comment.GetById() error", err)
		return nil, err
	}

	user, err := model.GetUserInfoById(comment.UserId)
	if err != nil {
		log.Error("ParentComment GetUserInfoById error", err)
		return nil, err
	}
	userInfo := model.UserHoleResponse2{
		Username:  user.Username,
		Avatar:    user.Avatar,
		Sid:       user.Sid,
		IsTeacher: user.IsTeacher,
	}
	//fmt.Println(comment.Time)

	data := &model.ParentCommentInfo{
		Id:             comment.Id,
		Content:        comment.Content,
		LikeNum:        0,
		IsLike:         false,
		Time:           comment.Time,
		UserInfo:       userInfo,
		SubCommentsNum: 0,
	}

	return data, nil
}
