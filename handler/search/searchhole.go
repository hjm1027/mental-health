package search

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/service"

	"github.com/gin-gonic/gin"
)

type searchHoleResponse struct {
	Sum   int                      `json:"sum"`
	Holes []model.HoleInfoResponse `json:"holes"`
}

func SearchHole(c *gin.Context) {
	userId := c.MustGet("id").(uint32)

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	keyword := c.DefaultQuery("keyword", "")
	tstr := c.DefaultQuery("type", "0")
	t, err := strconv.ParseUint(tstr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	holeList, err := service.GetAllHoles(keyword, page, limit, uint8(t), userId)
	if err != nil {
		handler.SendError(c, errno.ErrSearchHole, nil, err.Error())
		return
	}

	response := searchHoleResponse{
		Sum:   len(holeList),
		Holes: holeList,
	}
	handler.SendResponse(c, nil, response)
}
