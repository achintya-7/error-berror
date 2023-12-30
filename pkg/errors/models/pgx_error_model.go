package models

import "errors"

type PGXError struct {
	ErrMsg string
}

func (e PGXError) Error() string {
	return "[PGX ERROR] : " + e.ErrMsg
}

func (e PGXError) UnWrap() error {
	return errors.New(e.ErrMsg)
}
