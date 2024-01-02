package models

import "errors"

type MongoError struct {
	ErrMsg string
}

func (e MongoError) Error() string {
	return "[MONGO ERROR] : " + e.ErrMsg
}

func (e MongoError) UnWrap() error {
	return errors.New(e.ErrMsg)
}
