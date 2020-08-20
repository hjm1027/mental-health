package poster

import (
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

func GetPosterInfo(c *gin.Context) {
	//userId := c.MustGet("id").(uint32)

	data, err := model.LastestPoster()
	if err != nil {
		handler.SendError(c, errno.ErrGetPoster, nil, err.Error())
		return
	}

	response := model.PosterResponse{
		Home:     data.Home,
		Platform: data.Platform,
		Hole:     data.Hole,
	}

	handler.SendResponse(c, nil, response)
}
