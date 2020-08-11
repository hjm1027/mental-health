package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mental-health/model"

	"github.com/lexkong/log"
)

// 系统用户
var SystemUserId uint32 = 1

func MessageList(page, limit, uid uint32) (*[]model.MessageSub, error) {
	messages, err := model.GetMessages(page, limit, uid)
	if err != nil {
		return nil, nil
	}
	var messageSubs []model.MessageSub
	for _, message := range *messages {
		user := &model.UserModel{Id: message.PubUserId}
		err := user.GetUserById()
		if err != nil {
			return nil, err
		}
		userR := model.UserHoleResponse2{
			Sid:       user.Sid,
			IsTeacher: user.IsTeacher,
			Avatar:    user.Avatar,
			Username:  user.Username,
		}

		messageSub := model.MessageSub{
			UserInfo:        userR,
			Kind:            message.Kind,
			IsRead:          message.IsRead,
			Reply:           message.Reply,
			Time:            message.Time,
			HoleId:          message.HoleId,
			Content:         message.Content,
			Sid:             message.Sid,
			ParentCommentId: message.ParentCommentId,
		}

		messageSubs = append(messageSubs, messageSub)
	}
	return &messageSubs, nil
}

// 作出一级评论时，建立新的消息提醒
// 传的sid应为本用户的sid
func NewMessageForParentComment(userId uint32, comment *model.ParentCommentModel, hole *model.HoleModel) error {
	user := model.UserModel{Id: userId}
	if err := user.GetUserById(); err != nil {
		log.Info("user.GetUserById()  error")
		return err
	}

	message := &model.MessagePub{
		PubUserId:       userId,
		SubUserId:       hole.UserId,
		Kind:            2,
		IsRead:          false,
		Reply:           comment.Content,
		Time:            strconv.FormatInt(comment.Time.Unix(), 10),
		HoleId:          hole.Id,
		Content:         hole.Content,
		Sid:             user.Sid,
		ParentCommentId: comment.Id,
	}

	err := model.CreateMessage(message)
	if err != nil {
		log.Info("CreateMessage function error")
		return err
	}
	return nil
}

// 作出二级评论（回复）时，建立新的消息提醒
// 传的sid应为本用户的sid
func NewMessageForSubComment(userId uint32, comment *model.SubCommentModel, parentComment *model.ParentCommentModel) error {
	user := model.UserModel{Id: userId}
	if err := user.GetUserById(); err != nil {
		log.Info("user.GetUserById()  error")
		return err
	}

	hole := &model.HoleModel{Id: parentComment.HoleId}
	if err := hole.GetById(); err != nil {
		fmt.Println(parentComment.HoleId)
		log.Info("hole.GetById function error")
		return err
	}

	message := &model.MessagePub{
		PubUserId:       userId,
		SubUserId:       parentComment.UserId,
		Kind:            2,
		IsRead:          false,
		Reply:           comment.Content,
		Time:            strconv.FormatInt(comment.Time.Unix(), 10),
		HoleId:          parentComment.HoleId,
		Content:         parentComment.Content,
		Sid:             user.Sid,
		ParentCommentId: parentComment.Id,
	}

	err := model.CreateMessage(message)
	if err != nil {
		log.Info("CreateMessage function error")
		return err
	}
	return nil
}

func NewMessageForHoleLiking(userId uint32, hole *model.HoleModel) error {
	message := &model.MessagePub{
		PubUserId:       userId,
		SubUserId:       hole.UserId,
		Kind:            0,
		IsRead:          false,
		Reply:           "",
		Time:            strconv.FormatInt(time.Now().Unix(), 10),
		HoleId:          hole.Id,
		Content:         hole.Content,
		Sid:             "",
		ParentCommentId: 0,
	}

	err := model.CreateMessage(message)
	if err != nil {
		log.Info("CreateMessage function error")
		return err
	}
	return nil
}

func NewMessageForHoleFavoriting(userId uint32, hole *model.HoleModel) error {
	message := &model.MessagePub{
		PubUserId:       userId,
		SubUserId:       hole.UserId,
		Kind:            1,
		IsRead:          false,
		Reply:           "",
		Time:            strconv.FormatInt(time.Now().Unix(), 10),
		HoleId:          hole.Id,
		Content:         hole.Content,
		Sid:             "",
		ParentCommentId: 0,
	}

	err := model.CreateMessage(message)
	if err != nil {
		log.Info("CreateMessage function error")
		return err
	}
	return nil
}

func NewMessageForCommentLiking(userId, commentId uint32) error {
	comment, ok := model.IsSubComment(commentId)
	if ok {
		return NewMessageForSubCommentLiking(userId, comment)
	}
	return NewMessageForParentCommentLiking(userId, commentId)
}

func NewMessageForParentCommentLiking(userId, commentId uint32) error {
	comment := &model.ParentCommentModel{Id: commentId}
	if err := comment.GetById(); err != nil {
		log.Info("comment.GetById function error")
		return err
	}

	hole := &model.HoleModel{Id: comment.HoleId}
	if err := hole.GetById(); err != nil {
		log.Info("hole.GetById function error")
		return err
	}

	message := &model.MessagePub{
		PubUserId:       userId,
		SubUserId:       comment.UserId,
		Kind:            0,
		IsRead:          false,
		Reply:           "",
		Time:            strconv.FormatInt(time.Now().Unix(), 10),
		HoleId:          hole.Id,
		Content:         comment.Content,
		Sid:             "",
		ParentCommentId: 0,
	}

	err := model.CreateMessage(message)
	if err != nil {
		log.Info("CreateMessage function error")
		return err
	}
	return nil
}

func NewMessageForSubCommentLiking(userId uint32, comment *model.SubCommentModel) error {
	parentComment := &model.ParentCommentModel{Id: comment.Id}
	if err := parentComment.GetById(); err != nil {
		log.Info("parentComment.GetById function error")
		return err
	}

	hole := &model.HoleModel{Id: parentComment.HoleId}
	if err := hole.GetById(); err != nil {
		log.Info("hole.GetById function error")
		return err
	}

	message := &model.MessagePub{
		PubUserId:       userId,
		SubUserId:       comment.UserId,
		Kind:            0, //点赞
		IsRead:          false,
		Reply:           "",
		Time:            strconv.FormatInt(time.Now().Unix(), 10),
		HoleId:          hole.Id,
		Content:         comment.Content,
		Sid:             "",
		ParentCommentId: 0,
	}

	err := model.CreateMessage(message)
	if err != nil {
		log.Info("CreateMessage function error")
		return err
	}
	return nil
}
