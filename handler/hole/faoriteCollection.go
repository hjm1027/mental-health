package hole

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/service"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

type FavoriteCollectionResponse struct {
	Sum  int                       `json:"sum"`
	List []*model.HoleInfoResponse `json:"list"`
}

func GetFavoriteCollection(c *gin.Context) {
	userId := c.MustGet("id").(uint32)

	//不写limit就返回全部
	limitStr := c.DefaultQuery("limit", "1000")
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetQuery, nil, err.Error())
		return
	}

	data, err := service.GetCollectionList(userId, uint32(limit), uint32(page))
	if err != nil {
		log.Error("GetFavoriteCollection function error", err)
		handler.SendError(c, err, nil, err.Error())
	}

	handler.SendResponse(c, nil, FavoriteCollectionResponse{
		Sum:  len(data),
		List: data,
	})
}
