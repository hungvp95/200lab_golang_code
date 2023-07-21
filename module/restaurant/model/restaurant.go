package restaurantmodel

import (
	"errors"
	"food-delivery-200lab/common"
	"strings"
)

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Address         string `json:"address" gorm:"column:addr;"`
	Type            string `json:"type" gorm:"column:type;"`
}

func (Restaurant) TableName() string {
	return "restaurants_200lab"
}

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
	Status  *int    `json:"status" gorm:"column:status;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Address         string `json:"address" gorm:"column:addr;"`
}

func (data *RestaurantCreate) ValidateInputData() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrorNameIsEmpty
	}

	return nil
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

var (
	ErrorNameIsEmpty = errors.New("name can not empty")
)
