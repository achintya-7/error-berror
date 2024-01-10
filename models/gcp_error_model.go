package models

import "errors"

type GcpError struct {
	ErrMsg string
}

func (e GcpError) Error() string {
	return "[GCP ERROR] : " + e.ErrMsg
}

func (e GcpError) UnWrap() error {
	return errors.New(e.ErrMsg)
}

