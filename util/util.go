package util

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

// 获取当前时间，东八区
func GetCurrentTime() time.Time {
	// loc, _ := time.LoadLocation("Asia/Shanghai")
	// loc := time.FixedZone("CST", 8*3600)
	t := time.Now().UTC().Add(8 * time.Hour)
	return t
}

func ParseTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
