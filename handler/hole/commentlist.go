package hole

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/service"

	"github.com/gin-gonic/gin"
)

type commentListResponse struct {
	ParentCommentSum  int                         `json:"parent_comment_sum"`
	ParentCommentList *[]model.ParentCommentInfo2 `json:"parent_comment_list"`
}

func GetComments(c *gin.Context) {
	holeId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.ParseInt(limitStr, 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	pageNum := c.DefaultQuery("page", "1")
	page, err := strconv.ParseInt(pageNum, 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	userId := c.MustGet("id").(uint32)

	list, err := service.CommentList(uint32(holeId), int32(limit), int32((page-1)*limit), userId)
	if err != nil {
		handler.SendError(c, errno.ErrCommentList, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, commentListResponse{
		ParentCommentSum:  len(*list),
		ParentCommentList: list,
	})
}
