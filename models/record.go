package models

import "time"

type Record struct {
	Amount   int       `json:amount`
	DateTime time.Time `gorm:"column:DateTime" json:"datetime"`
}

type CreateRecordInput struct {
	Amount   int    `json:amount`
	DateTime string `json:"datetime"`
}

type StartEndTimeRecordInput struct {
	StartTime string `json:"StartDateTime"`
	EndTime   string `json:"EndDateTime"`
}
