package course

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type CourseInfoRequest struct {
	Url     string `json:"url"`
	Name    string `json:"name"`
	Source  string `json:"source"`
	Summary string `json:"summary"`
}

//修改课程信息
func PostInfo(c *gin.Context) {
	var data CourseInfoRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	//userId := c.MustGet("id").(uint32)

	courseId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	course := &model.CourseModel{
		Id:      uint32(courseId),
		Url:     data.Url,
		Name:    data.Name,
		Source:  data.Source,
		Summary: data.Summary,
	}
	if err := course.PostInfo(); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
