package service

import (
	"github.com/lexkong/log"

	"github.com/mental-health/model"
)

func GetCourseLikeCollection(userId, limit, page uint32) ([]*model.CourseInfoResponse, error) {
	var response []*model.CourseInfoResponse

	records, err := model.GetCourseLikeCollectionsByUserId(userId, limit, page)
	if err != nil {
		log.Error("GetCourseLikeCollectionsByUserId get records error", err)
		return nil, err
	}

	var courseIds []uint32
	for _, record := range *records {
		courseIds = append(courseIds, record.CourseId)
	}

	for i := 0; i < len(courseIds); i++ {
		course := model.CourseModel{Id: courseIds[i]}
		if err = course.GetInfo(); err != nil {
			log.Error("course.GetInfo() error", err)
			return nil, err
		}

		//get like state
		_, isLike := model.CourseHasLiked(userId, course.Id)
		_, isFavorite := model.CourseHasFavorited(userId, course.Id)

		data := &model.CourseInfoResponse{
			Id:          course.Id,
			Url:         course.Url,
			Name:        course.Name,
			Source:      course.Source,
			Summary:     course.Summary,
			LikeNum:     course.LikeNum,
			FavoriteNum: course.FavoriteNum,
			WatchNum:    course.WatchNum,
			IsLike:      isLike,
			IsFavorite:  isFavorite,
			Time:        course.Time,
		}

		response = append(response, data)
	}
	return response, nil
}
