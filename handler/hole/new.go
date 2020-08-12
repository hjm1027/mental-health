package hole

import (
	"unicode/utf8"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/util"

	"github.com/gin-gonic/gin"
)

type NewHoleRequest struct {
	Header  string `json:"header" binding:"required"`
	Type    uint8  `json:"type" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type NewHoleResponse struct {
	HoleId uint32 `json:"hole_id"`
}

func New(c *gin.Context) {
	var data NewHoleRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	userId := c.MustGet("id").(uint32)

	// Words are limited to 500
	// 字符数，非字节
	if utf8.RuneCountInString(data.Content) > 500 {
		handler.SendBadRequest(c, errno.ErrWordLimitation, nil, "Hole's content is limited to 500.")
		return
	}

	var hole = &model.HoleModel{
		UserId:      userId,
		Name:        data.Header,
		Content:     data.Content,
		LikeNum:     0,
		FavoriteNum: 0,
		CommentNum:  0,
		ReadNum:     0,
		Type:        data.Type,
		Time:        util.GetCurrentTime(),
	}
	//fmt.Println(hole.Time)

	if err := hole.New(); err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}
	response := NewHoleResponse{HoleId: hole.Id}

	handler.SendResponse(c, nil, response)
}
