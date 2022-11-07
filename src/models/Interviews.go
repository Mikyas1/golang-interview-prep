package models

import (
	"gorm.io/gorm"
	"time"
)

type InterviewStatus int

const (
	Ongoing InterviewStatus = iota
	UpComing
	Passed
	Failed
)

func (i InterviewStatus) String() string {
	return [...]string{
		"on_going",
		"up_coming",
		"passed",
		"failed",
	}[i]
}

type Interviews struct {
	gorm.Model
	User          User            `json:"user" gorm:"foreignKey:UserId;references:ID"`
	UserId        int             `json:"user_id" gorm:"not null"`
	Company       Company         `json:"company" gorm:"foreignKey:CompanyId;references:ID"`
	CompanyId     int             `json:"company_id" gorm:"not null"`
	InterviewDate time.Time       `json:"interview_date"`
	Status        InterviewStatus `json:"status"`
}
