package model

import (
	"errors"
	"time"
)

func (u *ReserveModel) TableName() string {
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
	if (data.Time.Year() == time.Year()) || (data.Time.YearDay() < time.YearDay()+int(data.AdvanceTime)-2) {
		return true, nil
	}
	if (data.Time.Year() < time.Year()) || (data.Time.YearDay() < time.YearDay()+leapYear(data.Time.Year())+int(data.AdvanceTime)-2) {
		return true, nil
	}
	return false, d.Error
}

func GetAllTeacher(limit, page uint32) (*[]UserModel, error) {
	var data []UserModel
	d := DB.Self.Table("user").Where("is_teacher = true").Limit(limit).Offset((page - 1) * limit).Find(&data)
	return &data, d.Error
}
