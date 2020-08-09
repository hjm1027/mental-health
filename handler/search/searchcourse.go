package search

import (
	"strconv"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/mental-health/service"

	"github.com/gin-gonic/gin"
)

type searchResponse struct {
	Sum     int                      `json:"sum"`
	Courses []model.CourseSearchInfo `json:"courses"`
}

func SearchCourse(c *gin.Context) {
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
	t := c.DefaultQuery("type", "0")
	//fmt.Println(t)
	tw := map[string]string{"0": "favorite_num DESC,like_num DESC,time DESC", "1": "time DESC", "2": "like_num DESC,time DESC", "3": "favorite_num DESC,time DESC"}
	//fmt.Println(tw[t])

	courseList := []model.CourseSearchInfo{}
	if keyword != "" {
		courseList, err = service.SearchCourses(keyword, page, limit, tw[t])
	} else {
		courseList, err = service.GetAllCourses(page, limit, tw[t])
	}
	if err != nil {
		handler.SendError(c, errno.ErrSearchCourse, nil, err.Error())
		return
	}

	response := searchResponse{
		Sum:     len(courseList),
		Courses: courseList,
	}
	handler.SendResponse(c, nil, response)
}
