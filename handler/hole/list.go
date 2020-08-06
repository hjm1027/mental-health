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

type HoleListResponse struct {
	Sum  int                       `json:"sum"`
	List []*model.HoleInfoResponse `json:"list"`
}

//树洞首页，获取问题列表
func GetHoleList(c *gin.Context) {
	userId := c.MustGet("id").(uint32)

	limitStr := c.DefaultQuery("limit", "10")
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

	data, err := service.GetHoleList(userId, uint32(limit), uint32(page))
	if err != nil {
		log.Error("GetHoleList function error", err)
		handler.SendError(c, err, nil, err.Error())
	}

	handler.SendResponse(c, nil, HoleListResponse{
		Sum:  len(data),
		List: data,
	})
}
