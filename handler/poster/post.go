package poster

import (
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 新增海报请求模型
type newPosterRequest struct {
	Home     string `json:"home" binding:"required"`
	Platform string `json:"platform" binding:"required"`
	Hole     string `json:"hole" binding:"required"`
}

func PostPosterInfo(c *gin.Context) {
	//userId := c.MustGet("id").(uint32)

	var data newPosterRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	poster := model.PosterModel{
		Home:     data.Home,
		Platform: data.Platform,
		Hole:     data.Hole,
	}

	if err := poster.New(); err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error())
	}

	handler.SendResponse(c, nil, nil)
}
