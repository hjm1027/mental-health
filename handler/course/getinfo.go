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

	response := model.CourseInfoResponse{
		Url:         course.Url,
		Name:        course.Name,
		Source:      course.Source,
		Summary:     course.Summary,
		LikeNum:     course.LikeNum,
		FavoriteNum: course.FavoriteNum,
		WatchNum:    course.WatchNum,
		Time:        course.Time,
	}
	handler.SendResponse(c, nil, response)
}
