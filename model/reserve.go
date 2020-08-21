package model

import (
	"errors"
	"time"
)

func (u *ReserveModel) TableName() string {
	return "reserve"
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
	if (data.Time.Year() < time.Year()) || (data.Time.YearDay() < time.YearDay()+2) {
		return true, nil
	}
	return false, d.Error
}
