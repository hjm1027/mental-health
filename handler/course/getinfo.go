package course

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

//获取课程信息
func GetInfo(c *gin.Context) {
	userId := c.MustGet("id").(uint32)

	courseId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	course := &model.CourseModel{Id: uint32(courseId)}
	if err := course.GetInfo(); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	_, isLike := model.CourseHasLiked(userId, course.Id)
	_, isFavorite := model.CourseHasFavorited(userId, course.Id)

	response := model.CourseInfoResponse{
		Id:          course.Id,
		Url:         course.Url,
		Name:        course.Name,
		Source:      course.Source,
		Summary:     course.Summary,
		LikeNum:     course.LikeNum,
		FavoriteNum: course.FavoriteNum,
		WatchNum:    course.WatchNum,
		Time:        course.Time,
		IsLike:      isLike,
		IsFavorite:  isFavorite,
	}
	handler.SendResponse(c, nil, response)
}
