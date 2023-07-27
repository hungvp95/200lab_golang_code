package uploadmodel

import (
	"errors"
	"food-delivery-200lab/common"
)

type Upload struct {
	common.Image `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileTooLarge",
	)
)

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err, "file is not image", "ErrFileIsNotImage")
}

func ErrCanNotSaveFile(err error) *common.AppError {
	return common.NewCustomError(err, "can not save uploaded file", "ErrCanNotSaveFile")
}
