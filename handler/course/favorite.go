package course

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type FavoriteDataRequest struct {
	FavoriteState bool `json:"favorite_state"`
}

func FavoriteCourse(c *gin.Context) {
	courseId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	userId := c.MustGet("id").(uint32)

	recordId, hasFavorited := model.CourseHasFavorited(userId, uint32(courseId))

	// 获取请求中当前的点赞状态
	var bodyData FavoriteDataRequest
	if err := c.BindJSON(&bodyData); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 未点赞
	if bodyData.FavoriteState && !hasFavorited {
		handler.SendResponse(c, errno.ErrNotFavorited, nil)
		return
	}
	// 已点赞
	if !bodyData.FavoriteState && hasFavorited {
		handler.SendResponse(c, errno.ErrHasFavorited, nil)
		return
	}

	// 点赞或者取消点赞
	if bodyData.FavoriteState {
		err = model.CourseUnfavorite(recordId)
	} else {
		err = model.CourseFavorite(userId, uint32(courseId))
	}

	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
