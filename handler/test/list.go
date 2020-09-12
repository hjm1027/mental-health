package test

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

type TestListResponse struct {
	Sum  int                `json:"sum"`
	List []*model.TestModel `json:"list"`
}

func GetList(c *gin.Context) {
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

	testList, err := model.GetList(uint32(limit), uint32(page))
	if err != nil {
		handler.SendError(c, errno.ErrGetTestList, nil, err.Error())
		return
	}

	response := TestListResponse{
		Sum:  len(testList),
		List: testList,
	}

	handler.SendResponse(c, nil, response)
}
