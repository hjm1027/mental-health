package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/mental-health/handler"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/service"
)

type UploadImageResponse struct {
	Url string `json:"url"`
}

func Image(c *gin.Context) {
	image, header, err := c.Request.FormFile("image")
	if err != nil {
		handler.SendError(c, errno.ErrGetFile, nil, err.Error())
		return
	}
	dataLen := header.Size
	userId := c.MustGet("id").(uint32)

	url, err := service.UploadImage(header.Filename, userId, image, dataLen)
	if err != nil {
		handler.SendError(c, errno.ErrUploadFile, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, UploadImageResponse{Url: url})
}
