package hole

import (
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/lexkong/log"
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/service"
	"github.com/mental-health/util"

	"github.com/gin-gonic/gin"
)

type commentReplyResponse struct {
	Id             uint32                  `json:"id"` // 回复的评论的id
	Content        string                  `json:"content"`
	LikeNum        uint32                  `json:"like_num"`
	IsLike         bool                    `json:"is_like"`
	Time           time.Time               `json:"time"`
	UserInfo       model.UserHoleResponse2 `json:"user_info"`
	TargetUserInfo model.TargetUserInfo    `json:"target_user_info"`
}

func Reply(c *gin.Context) {
	var data newCommentRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userId := c.MustGet("id").(uint32)
	parentId, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	// Get the user's sid whom is reply to.
	sid, ok := c.GetQuery("sid")
	if !ok {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "Target-user's sid is expected.")
		return
	}

	userinfo := &model.UserModel{Sid: sid}
	if err := userinfo.GetUserBySid(); err != nil {
		log.Error("userinfo.GetUserBySid() error", err)
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "Target-user's sid is error.")
		return
	}

	// Get parentComment by id
	parentComment := &model.ParentCommentModel{Id: uint32(parentId)}
	if err := parentComment.GetById(); err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, "The parent comment does not exist.")
		return
	}

	// Words are limited to 300
	if utf8.RuneCountInString(data.Content) > 300 {
		handler.SendBadRequest(c, errno.ErrWordLimitation, nil, "Comment's content is limited to 300.")
		return
	}

	var comment = &model.SubCommentModel{
		UserId:       userId,
		ParentId:     uint32(parentId),
		TargetUserId: parentComment.Id,
		Content:      data.Content,
		Time:         util.GetCurrentTime(),
	}

	// Create a new subComment
	commentId, err := comment.New()
	if err != nil {
		log.Error("comment.New function error. ", err)
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	comment.Id = commentId

	// Add one to the parentComment's subCommentNum
	if err := parentComment.UpdateSubCommentNum(1); err != nil {
		log.Error("UpdateSubCommentNum function error. ", err)
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	user := model.UserModel{Id: userId}
	err = user.GetUserById()
	if err != nil {
		log.Error("reply GetUserById error", err)
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	userinfo2 := model.UserHoleResponse2{
		Username:  user.Username,
		Avatar:    user.Avatar,
		Sid:       user.Sid,
		IsTeacher: user.IsTeacher,
	}

	targetinfo := model.TargetUserInfo{
		Id: userinfo.Id,
	}

	response := &commentReplyResponse{
		Id:             commentId,
		Content:        data.Content,
		LikeNum:        0,
		IsLike:         false,
		Time:           util.GetCurrentTime(),
		UserInfo:       userinfo2,
		TargetUserInfo: targetinfo,
	}

	handler.SendResponse(c, nil, response)

	err = service.NewMessageForSubComment(userId, comment, parentComment)
	if err != nil {
		log.Error("NewMessageForSubComment failed", err)
	}
}
