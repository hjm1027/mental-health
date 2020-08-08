package course

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/service"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

type LikeCollectionResponse struct {
	Sum  int                         `json:"sum"`
	List []*model.CourseInfoResponse `json:"list"`
}

func GetLikeCollection(c *gin.Context) {
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

	data, err := service.GetCourseLikeCollection(userId, uint32(limit), uint32(page))
	if err != nil {
		log.Error("GetLikeCollection function error", err)
		handler.SendError(c, err, nil, err.Error())
	}

	handler.SendResponse(c, nil, LikeCollectionResponse{
		Sum:  len(data),
		List: data,
	})
}
