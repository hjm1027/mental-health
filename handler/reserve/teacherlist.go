package reserve

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

//获取老师信息列表
func TeacherList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	teacherList, err := model.GetAllTeacher(uint32(limit), uint32(page))
	if err != nil {
		handler.SendError(c, errno.ErrGetTeacherList, nil, err.Error())
		return
	}

	handler.SendResponse(c, errno.OK, teacherList)
}
