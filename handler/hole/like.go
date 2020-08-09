package hole

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type LikeDataRequest struct {
	LikeState bool `json:"like_state"`
}

func LikeHole(c *gin.Context) {
	holeId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	//get hole
	hole := &model.HoleModel{Id: uint32(holeId)}
	if err = hole.GetById(); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	userId := c.MustGet("id").(uint32)

	recordId, hasLiked := model.HasLiked(userId, uint32(holeId))

	// 获取请求中当前的点赞状态
	var bodyData LikeDataRequest
	if err := c.BindJSON(&bodyData); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 未点赞
	if bodyData.LikeState && !hasLiked {
		handler.SendResponse(c, errno.ErrNotLiked, nil)
		return
	}
	// 已点赞
	if !bodyData.LikeState && hasLiked {
		handler.SendResponse(c, errno.ErrHasLiked, nil)
		return
	}

	var err2 error

	// 点赞或者取消点赞
	if bodyData.LikeState {
		err = model.Unlike(recordId)
		if err != nil {
			handler.SendError(c, errno.ErrDatabase, nil, err.Error())
			return
		}
		err2 = hole.UpdateLikeNum(-1)
	} else {
		err = model.Like(userId, uint32(holeId))
		if err != nil {
			handler.SendError(c, errno.ErrDatabase, nil, err.Error())
			return
		}
		err2 = hole.UpdateLikeNum(1)
	}

	if err2 != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err2.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
