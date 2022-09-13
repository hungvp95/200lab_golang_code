package model

import (
	"errors"
	"strings"
)


type RestaurantCreate struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (r *RestaurantCreate) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.New("restaurant name can't be blank")
	}
	r.Addr = strings.TrimSpace(r.Addr)
	if r.Addr == "" {
		return errors.New("restaurant address can't be blank")
	}

	return nil
}