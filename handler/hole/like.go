package hole

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type likeDataRequest struct {
	LikeState bool `json:"like_state"`
}

func LikeHole(c *gin.Context) {
	holeId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	userId := c.MustGet("id").(uint32)

	recordId, hasLiked := model.HasLiked(userId, uint32(holeId))

	// 获取请求中当前的收藏状态
	var bodyData likeDataRequest
	if err := c.BindJSON(&bodyData); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 未收藏
	if bodyData.LikeState && !hasLiked {
		handler.SendResponse(c, errno.ErrNotLiked, nil)
		return
	}
	// 已收藏
	if !bodyData.LikeState && hasLiked {
		handler.SendResponse(c, errno.ErrHasLiked, nil)
		return
	}

	// 收藏或者取消收藏
	if bodyData.LikeState {
		err = model.Unlike(recordId)
	} else {
		err = model.Like(userId, uint32(holeId))
	}

	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
