package service

import (
	"github.com/lexkong/log"
	"github.com/mental-health/model"
)

func SearchCourses(keyword string, page, limit uint64, t string) ([]model.CourseSearchInfo, error) {
	courseRows, _ := model.AgainstAndMatchCourses(keyword, page, limit, t)
	courses := make([]model.CourseSearchInfo, len(courseRows))
	for i, row := range courseRows {
		courses[i] = model.CourseSearchInfo{
			Id:          row.Id,
			Url:         row.Url,
			Name:        row.Name,
			Source:      row.Source,
			Summary:     row.Summary,
			LikeNum:     row.LikeNum,
			FavoriteNum: row.FavoriteNum,
		}
	}
	return courses, nil
}

func GetAllCourses(page, limit uint64, t string) ([]model.CourseSearchInfo, error) {
	courseRows, err := model.AllCourses(page, limit, t)
	if err != nil {
		return nil, err
	}
	courses := make([]model.CourseSearchInfo, len(courseRows))
	for i, row := range courseRows {
		courses[i] = model.CourseSearchInfo{
			Id:          row.Id,
			Url:         row.Url,
			Name:        row.Name,
			Source:      row.Source,
			Summary:     row.Summary,
			LikeNum:     row.LikeNum,
			FavoriteNum: row.FavoriteNum,
		}
	}
	return courses, nil
}

func GetAllHoles(keyword string, page, limit uint64, t uint8, userId uint32) ([]model.HoleInfoResponse, error) {
	var err error
	var holeRows []model.HoleModel
	if keyword == "" {
		holeRows, err = model.AllHoles(page, limit, t)
	} else {
		holeRows, err = model.SearchHoles(keyword, page, limit, t)
	}
	if err != nil {
		return nil, err
	}
	holes := make([]model.HoleInfoResponse, len(holeRows))
	for i, row := range holeRows {
		//get user
		user, err := model.GetUserInfoById(userId)
		if err != nil {
			log.Error("GetAllHole get user error", err)
		}
		userInfo := model.UserHoleResponse{
			Username: user.Username,
			Avatar:   user.Avatar,
			Sid:      user.Sid,
		}

		//get like state
		_, isLike := row.HasLiked(userId)
		_, isFavorite := row.HasFavorited(userId)

		holes[i] = model.HoleInfoResponse{
			HoleId:      row.Id,
			Type:        row.Type,
			Name:        row.Name,
			Content:     row.Content,
			LikeNum:     row.LikeNum,
			ReadNum:     row.ReadNum,
			FavoriteNum: row.FavoriteNum,
			IsLike:      isLike,
			IsFavorite:  isFavorite,
			Time:        row.Time,
			CommentNum:  row.CommentNum,
			UserInfo:    userInfo,
		}
	}
	return holes, nil
}
