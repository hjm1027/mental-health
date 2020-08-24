package reserve

import (
	"net/smtp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mental-health/handler"
	"github.com/mental-health/model"
	"github.com/mental-health/pkg/errno"
	"github.com/spf13/viper"
)

type ReserveRequest struct {
	Weekday  uint8  `json:"weekday" binding:"required"`
	Schedule uint8  `json:"schedule" binding:"required"`
	Type     uint8  `json:"type" binding:"required"`
	Method   uint8  `json:"method"`
	ModuleId uint32 `json:"module_id"`
}

func getAdvanceTime(weekday, weekdayNow uint8) uint8 {
	difference := map[int]int{-6: 8, -5: 2, -4: 3, -3: 4, -2: 5, -1: 6, 0: 7, 1: 8, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
	advance := difference[int(weekday)-int(weekdayNow)]
	return uint8(advance)
}

func sendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])

	msg := []byte("To:" + to + "\nFrom: " + user + "<" +
		user + ">\nSubject: " + subject + "\n" + "\n\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func sendEmail(subject, body string) error {
	user := viper.GetString("mail.user")
	pwd := viper.GetString("mail.password")
	host := viper.GetString("mail.host")
	to := viper.GetString("mail.to")
	err := sendToMail(user, pwd, host, to, subject, body, "html")
	return err
}

func sendEmail2(subject, body string) error {
	user := viper.GetString("mail.user2")
	pwd := viper.GetString("mail.password2")
	host := viper.GetString("mail.host")
	to := viper.GetString("mail.to2")
	err := sendToMail(user, pwd, host, to, subject, body, "html")
	return err
}

func parsew(weekday uint8) string {
	w := map[uint8]string{1: "星期一", 2: "星期二", 3: "星期三", 4: "星期四", 5: "星期五", 6: "星期六", 7: "星期日"}
	return w[weekday]
}

func parses(schedule uint8) string {
	s := map[uint8]string{1: "8:10-9:00", 2: "9:10-10:00", 3: "10:20-11:10", 4: "14:10-15:00", 5: "15:10-16:00", 6: "16:20-17:10"}
	return s[schedule]
}

func parset(Type uint8) string {
	t := map[uint8]string{1: "环境适应类", 2: "人际关系类", 3: "学业学习类", 4: "生活经济类", 5: "求职择业类", 6: "其他"}
	return t[Type]
}

func parsem(method uint8) string {
	m := map[uint8]string{0: "线上预约", 1: "线下预约"}
	return m[method]
}

//进行预约
func Reserve(c *gin.Context) {
	userId := c.MustGet("id").(uint32)
	user := model.UserModel{Id: userId}
	if err := user.GetUserById(); err != nil {
		handler.SendError(c, errno.ErrGetUserInfo, nil, err.Error())
	}

	var data ReserveRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		handler.SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	time2 := time.Now().UTC().Add(8 * time.Hour)
	weekday := time.Now().UTC().Add(8 * time.Hour).Weekday()

	teacher, err := model.GetTeacherBySchedule(data.Weekday, data.Schedule)
	if err != nil {
		handler.SendError(c, errno.ErrGetTeacherBySchedule, nil, err.Error())
		return
	}

	advanceTime := getAdvanceTime(data.Weekday, uint8(weekday))

	reserve := &model.ReserveModel{
		Weekday:     data.Weekday,
		Schedule:    data.Schedule,
		Teacher:     teacher,
		Reserve:     1,
		Time:        time2,
		AdvanceTime: advanceTime,
		Type:        data.Type,
		Method:      data.Method,
		UserId:      userId,
	}

	err1, err2 := reserve.New(userId)
	if err1 != nil {
		handler.SendError(c, errno.ErrCreateReserve, nil, err1.Error())
		return
	}
	if err2 != nil {
		handler.SendError(c, errno.ErrCreateReserve, nil, err2.Error())
		return
	}

	//w := parsew(data.Weekday)
	//s := parses(data.Schedule)
	//t := parset(data.Type)

	subject := "心理预约请求"
	body := "尊敬的" + teacher + ",学号为 " + user.Sid + " 的学生向您发出预约申请，预约时间为 " + parsew(data.Weekday) + "的" + parses(data.Schedule) + ",咨询类型为 " + parset(data.Type) + "咨询方式为 " + parsem(data.Method) + "。您可以同意或者拒绝。（随便写的文案，以后再改）"

	err = sendEmail(subject, body)
	if err != nil {
		handler.SendError(c, errno.ErrSendMail, nil, err.Error())
		return
	}

	timer := time.NewTimer(30 * time.Second)

	go func() {
		select {
		case <-timer.C:
			status, err := reserve.Status()
			if err != nil {
				handler.SendError(c, errno.ErrGetStatus, nil, err.Error())
				return
			}

			if status == 1 {
				subject2 := "这个老师没有回复"
				body2 := "赶快提醒他，老师姓名为 " + teacher + "。"
				err = sendEmail2(subject2, body2)
				if err != nil {
					handler.SendError(c, errno.ErrSendMail, nil, err.Error())
					return
				}
			}
		}
	}()
	handler.SendResponse(c, errno.OK, nil)
}
