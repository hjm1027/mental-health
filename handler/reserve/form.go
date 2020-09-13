package reserve

import (
	"time"

	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"

	"github.com/gin-gonic/gin"
)

/*
	Schedule1 uint8 `json:"schedule1"`
	Teacher1 uint32 `json:"teacher1"`
	Schedule2 uint8 `json:"schedule2"`
	Teacher2 uint32 `json:"teacher2"`
	Schedule3 uint8 `json:"schedule3"`
	Teacher3 uint32 `json:"teacher3"`
	Schedule4 uint8 `json:"schedule4"`
	Teacher4 uint32 `json:"teacher4"`
	Schedule5 uint8 `json:"schedule5"`
	Teacher5 uint32 `json:"teacher5"`
	Schedule6 uint8 `json:"schedule6"`
	Teacher6 uint32 `json:"teacher6"`*/

type FormResponse struct {
	Weekday uint8      `json:"weekday"`
	Date    string     `json:"date"`
	Item    []FormItem `json:"item"`
}

type FormItem struct {
	Id      uint8  `json:"id"`
	Teacher uint32 `json:"teacher"`
	Reserve uint8  `json:"reserve"`
}

func fix(d uint8) uint8 {
	if d <= 7 {
		return d
	} else {
		return d - 7
	}
}

func fix2(w int) int {
	if w > 0 {
		return w
	} else {
		return 7
	}
}

//获取预约状态表格
func ReserveForm(c *gin.Context) {
	var response []FormResponse

	time1 := time.Now().UTC().Add(8 * time.Hour)
	time2 := time.Now().UTC().Add(56 * time.Hour)
	weekday := fix2(int(time2.Weekday()))
	//fmt.Println(weekday)
	var n int

	for d := uint8(weekday); d < uint8(weekday)+7; d++ {
		var item []FormItem
		for i := 1; i <= 6; i++ {
			reserve, err := model.GetReserveBySchedule(fix(d), uint8(i))
			if err != nil {
				handler.SendError(c, errno.ErrGetReserve, nil, err.Error())
				return
			}

			var status uint8
			if reserve.Reserve == 0 || reserve.Reserve == 1 {
				status = reserve.Reserve
			} else {
				status = model.QueryReserve2(reserve, time1)
			}

			formItem := FormItem{
				Id:      uint8(i),
				Teacher: reserve.TeacherId,
				Reserve: status,
			}

			item = append(item, formItem)
		}

		date := time.Now().UTC().Add(56 * time.Hour).Add(time.Duration(24*n) * time.Hour)
		//fmt.Println(date)
		dateFix := date.Format("2006-01-02 15:04:05")
		//fmt.Println(date.Month())
		//fmt.Println(date.Day())
		n++
		//month := string(date.Month())

		formResponse := FormResponse{
			Weekday: fix(d),
			Date:    dateFix[:10],
			Item:    item,
		}

		response = append(response, formResponse)
	}

	handler.SendResponse(c, errno.OK, response)
}
