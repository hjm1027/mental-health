package hole

import (
	"strconv"

	"github.com/lexkong/log"
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type CommentLikeRequest struct {
	LikeState bool `json:"like_state"`
}

type CommentLikeResponse struct {
	LikeState bool   `json:"like_state"`
	LikeNum   uint32 `json:"like_num"`
}

func CommentLike(c *gin.Context) {
	var err error
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	// 获取请求中的点赞状态
	var bodyData CommentLikeRequest
	if err := c.BindJSON(&bodyData); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userId := c.MustGet("id").(uint32)

	likeRecordId, hasLiked := model.CommentHasLiked(userId, uint32(id))

	// 判断点赞请求是否合理
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

	// 点赞&取消点赞
	if bodyData.LikeState {
		err = model.CommentCancelLiking(likeRecordId)
	} else {
		err = model.CommentLiking(userId, uint32(id))
	}

	if err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}

	likeNum, err := model.GetCommentLikeSum(uint32(id))
	if err != nil {
		log.Error("model.GetCommentLikeSum() error", err)
		handler.SendError(c, err, nil, err.Error())
		return
	}
	//fmt.Println(likeNum)
	data := &CommentLikeResponse{
		LikeState: !hasLiked,
		LikeNum:   likeNum,
	}

	handler.SendResponse(c, nil, data)
}
