package reserve

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type RecordListResponse struct {
	Sum  int                  `json:"sum"`
	List []*model.RecordModel `json:"list"`
}

//获取我的记录
func GetRecord(c *gin.Context) {
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

	records, err := model.GetRecords(userId, uint32(page), uint32(limit))
	if err != nil {
		handler.SendError(c, errno.ErrGetRecord, nil, err.Error())
		return
	}

	response := RecordListResponse{
		Sum:  len(records),
		List: records,
	}

	handler.SendResponse(c, errno.OK, response)
}
