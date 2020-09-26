package hole

import (
	"strconv"
	"unicode/utf8"

	"github.com/lexkong/log"
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/service"
	"github.com/mental-health/util"
	"github.com/mental-health/util/security"

	"github.com/gin-gonic/gin"
)

// 新增评论请求模型
type newCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

func NewParentComment(c *gin.Context) {
	userId := c.MustGet("id").(uint32)

	var data newCommentRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	holeId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	// Words are limited to 300
	if utf8.RuneCountInString(data.Content) > 300 {
		handler.SendBadRequest(c, errno.ErrWordLimitation, nil, "Comment's content is limited to 300.")
		return
	}

	// 小程序内容安全检测
	ok, err := security.MsgSecCheck(data.Content)
	if err != nil {
		handler.SendError(c, errno.ErrSecurityCheck, nil, err.Error())
		return
	} else if !ok {
		log.Errorf(err, "WX security check msg(%s) error", data.Content)
		handler.SendBadRequest(c, errno.ErrSecurityCheck, nil, "hole comment content violation")
		return
	}

	var comment = &model.ParentCommentModel{
		UserId:        userId,
		HoleId:        uint32(holeId),
		Content:       data.Content,
		Time:          util.GetCurrentTime(),
		SubCommentNum: 0,
	}

	// Create new comment
	commentId, err := comment.New()
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}
	//commentId, _ := strconv.Atoi(commentIdstr)

	//fmt.Println(comment.Id)

	// Add one to the hole's comment sum
	hole := &model.HoleModel{Id: uint32(holeId)}
	if err := hole.GetById(); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	if err := hole.UpdateCommentNum(); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	// Get comment info
	commentInfo, err := service.GetParentCommentInfo(uint32(commentId), userId)
	if err != nil {
		handler.SendError(c, errno.ErrGetParentCommentInfo, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, commentInfo)

	err = service.NewMessageForParentComment(userId, comment, hole)
	if err != nil {
		log.Error("NewMessageForParentComment failed", err)
	}
}
