package model

import (
	"github.com/definev/200lab_golang/food_delivery/common"
)

var (
	EntityName = "Restaurant"
)

type Restaurant struct {
	common.SqlModel `json:",inline"`
	Name            string `json:"name" gorm:"not null;column:name;"`
	Addr            string `json:"addr" gorm:"not null;column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}
