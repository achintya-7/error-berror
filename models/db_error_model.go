package models

import "errors"

type DBError struct {
	ErrMsg string
}

func (e DBError) Error() string {
	return "[DB ERROR] : " + e.ErrMsg
}

func (e DBError) UnWrap() error {
	return errors.New(e.ErrMsg)
}
