package hole

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

func FavoriteHole(c *gin.Context) {
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

	recordId, hasFavorited := model.HasFavorited(userId, uint32(holeId))

	// 获取请求中当前的收藏状态
	var bodyData FavoriteDataRequest
	if err := c.BindJSON(&bodyData); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	// 未收藏
	if bodyData.FavoriteState && !hasFavorited {
		handler.SendResponse(c, errno.ErrNotFavorited, nil)
		return
	}
	// 已收藏
	if !bodyData.FavoriteState && hasFavorited {
		handler.SendResponse(c, errno.ErrHasFavorited, nil)
		return
	}

	var err2 error

	// 收藏或者取消收藏
	if bodyData.FavoriteState {
		err = model.Unfavorite(recordId)
		if err != nil {
			handler.SendError(c, errno.ErrDatabase, nil, err.Error())
			return
		}
		err2 = hole.UpdateFavoriteNum(-1)
	} else {
		err = model.Favorite(userId, uint32(holeId))
		if err != nil {
			handler.SendError(c, errno.ErrDatabase, nil, err.Error())
			return
		}
		err2 = hole.UpdateFavoriteNum(1)
	}

	if err2 != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err2.Error())
		return
	}

	handler.SendResponse(c, nil, "ok")
}
