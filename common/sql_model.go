package common

import "time"

type SQLModel struct {
	Id       int        `json:"id" gorm:"column:id;"`
	Status   int        `json:"status" gorm:"column:status;default:1;"`
	CreateAt *time.Time `json:"createAt,omitempty" gorm:"column:created_at;"`
	UpdateAt *time.Time `json:"updateAt,omitempty" gorm:"column:updated_at;"`
}
