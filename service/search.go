package service

import (
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
