package hole

import (
	"strconv"
	"time"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type HoleInfoResponse struct {
	HoleId      uint32                 `json:"hole_id"`
	Type        uint8                  `json:"type"`
	Content     string                 `json:"content"`
	LikeNum     uint32                 `json:"like_num"`
	ReadNum     uint32                 `json:"read_num"`
	FavoriteNum uint32                 `json:"favorite_num"`
	IsLike      bool                   `json:"is_like"`
	IsFavorite  bool                   `json:"is_favorite"`
	Time        time.Time              `json:"time"`
	CommentNum  uint32                 `json:"comment_num"`
	UserInfo    model.UserHoleResponse `json:"user_info"`
}

func GetHoleInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetParam, nil, err.Error())
		return
	}

	userId := c.MustGet("id").(uint32)

	//get hole
	hole := model.HoleModel{Id: uint32(id)}
	if err := hole.GetById(); err != nil {
		handler.SendError(c, errno.ErrGetHoleInfo, nil, err.Error())
		return
	}

	//get user
	user, err := model.GetUserInfoById(userId)
	if err != nil {
		handler.SendError(c, errno.ErrGetUserInfo, nil, err.Error())
	}
	userInfo := model.UserHoleResponse{
		Username: user.Username,
		Avatar:   user.Avatar,
	}

	//read_num +1
	_, isRead := hole.HasRead(userId)
	newNum := hole.ReadNum
	if !isRead {
		newNum += 1
		//fmt.Println("reading")
		holeRead := model.HoleReadModel{
			HoleId: hole.Id,
			UserId: userId,
		}
		err := holeRead.Read()
		if err != nil {
			handler.SendError(c, errno.ErrGetUserInfo, nil, err.Error())
		}
	}

	//get like state
	_, isLike := hole.HasLiked(userId)
	_, isFavorite := hole.HasFavorited(userId)

	data := HoleInfoResponse{
		HoleId:      hole.Id,
		Type:        hole.Type,
		Content:     hole.Content,
		LikeNum:     hole.LikeNum,
		ReadNum:     newNum,
		FavoriteNum: hole.FavoriteNum,
		IsLike:      isLike,
		IsFavorite:  isFavorite,
		Time:        hole.Time,
		CommentNum:  hole.CommentNum,
		UserInfo:    userInfo,
	}

	handler.SendResponse(c, nil, data)
}
