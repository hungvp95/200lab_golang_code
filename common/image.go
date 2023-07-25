package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloudName,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
}

func (img *Image) Scan(value interface{}) error {
	bytes, isOk := value.([]byte)
	if !isOk {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var image Image
	if err := json.Unmarshal(bytes, &image); err != nil {
		return err
	}

	*img = image
	return nil
}

// Value return json value, implement driver.Value interface
func (img *Image) Value() (driver.Value, error) {
	if img == nil {
		return nil, nil
	}
	return json.Marshal(img)
}

type Images []Image

func (img *Images) Scan(value interface{}) error {
	bytes, isOk := value.([]byte)
	if !isOk {
		return errors.New(fmt.Sprint("Failed to unmarshal JSON value:", value))
	}

	var image []Image
	if err := json.Unmarshal(bytes, &image); err != nil {
		return err
	}

	*img = image
	return nil
}

// Value return json value, implement driver.Value interface
func (img *Images) Value() (driver.Value, error) {
	if img == nil {
		return nil, nil
	}
	return json.Marshal(img)
}
