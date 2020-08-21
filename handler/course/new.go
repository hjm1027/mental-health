package course

import (
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/util"

	"github.com/gin-gonic/gin"
)

func NewCourse(c *gin.Context) {
	url := c.Query("url")
	name := c.Query("name")
	source := c.Query("source")
	summary := c.Query("summary")
	if url == "" || name == "" || source == "" || summary == "" {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, "")
		return
	}

	course := model.CourseModel{
		Url:         url,
		Name:        name,
		Source:      source,
		Summary:     summary,
		LikeNum:     0,
		FavoriteNum: 0,
		WatchNum:    0,
		Time:        util.GetCurrentTime(),
	}

	id, err := course.New()
	if err != nil {
		handler.SendError(c, err, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, id)
}
