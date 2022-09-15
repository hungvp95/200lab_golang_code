package common

import "errors"

var (
	ErrRecordNotFound = errors.New("Record not found!")
	ErrStatusDifferentZero = errors.New("Status different zero")
)