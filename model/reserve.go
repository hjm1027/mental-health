package model

import (
	"errors"
	"time"
)

func (reserve *ReserveModel) TableName() string {
	return "reserve"
}

func leapYear(year int) int {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return 366
	}
	return 365
}

func QueryReserve(weekday, schedule uint8, time time.Time) (bool, error) {
	var data ReserveModel
	d := DB.Self.Table("reserve").Where("weekday = ? AND schedule = ?", weekday, schedule).First(&data)
	if d.RecordNotFound() {
		return false, errors.New("Out of range data weekday or schedule")
	}
	if data.Reserve == 0 {
		return true, nil
	}
	if data.Time.IsZero() {
		return true, nil
	}
	if (data.Time.Year() == time.Year()) && (data.Time.YearDay() < time.YearDay()-int(data.AdvanceTime)+2) {
		return true, nil
	}
	if (data.Time.Year() < time.Year()) && (data.Time.YearDay() < time.YearDay()+leapYear(data.Time.Year())-int(data.AdvanceTime)+2) {
		return true, nil
	}
	return false, d.Error
}

func GetAllTeacher(limit, page uint32) (*[]UserModel, error) {
	var data []UserModel
	d := DB.Self.Table("user").Where("is_teacher = true").Limit(limit).Offset((page - 1) * limit).Find(&data)
	return &data, d.Error
}

func GetTeacherBySchedule(weekday, schedule uint8) (string, error) {
	var reserve ReserveModel
	d := DB.Self.Table("reserve").Where("weekday = ? AND schedule = ?", weekday, schedule).First(&reserve)
	return reserve.Teacher, d.Error
}

func (reserve *ReserveModel) New(userId uint32) (error, error) {
	var data ReserveModel
	d := DB.Self.Table("reserve").Where("weekday = ? AND schedule = ?", reserve.Weekday, reserve.Schedule).First(&data)
	time := time.Now().UTC().Add(8 * time.Hour)
	data.Reserve = 1
	data.Time = time
	data.AdvanceTime = reserve.AdvanceTime
	data.Type = reserve.Type
	data.Method = reserve.Method
	data.UserId = userId
	d2 := DB.Self.Save(data)
	return d.Error, d2.Error
}

func (reserve *ReserveModel) Status() (uint8, error) {
	var data ReserveModel
	d := DB.Self.Table("reserve").Where("weekday = ? AND schedule = ?", reserve.Weekday, reserve.Schedule).First(&data)
	return data.Reserve, d.Error
}

func CheckReserve(weekday, schedule, status uint8) error {
	var data ReserveModel
	d := DB.Self.Table("reserve").Where("weekday = ? AND schedule = ?", weekday, schedule).First(&data)
	data.Reserve = status
	d = DB.Self.Save(data)
	return d.Error
}
