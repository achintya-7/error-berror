package errors

import (
	"errors"
)

type DBError struct {
	ErrMsg string
}

// Custom DB errors
var (
	ErrDBConnectionFailed = DBError{ErrMsg: "DB connection failed"}
	ErrDBQueryFailed      = DBError{ErrMsg: "DB query failed"}
)

func (e DBError) Error() string {
	return "[DB ERROR] : " + e.ErrMsg
}

func (e DBError) UnWrap() error {
	return errors.New(e.ErrMsg)
}

func HandleDBException(err error) (string, string) {
	switch err {
	case ErrDBConnectionFailed:
		return ErrDBConnectionFailed.Error(), ErrDBConnectionFailed.Error()

	case ErrDBQueryFailed:
		return ErrDBQueryFailed.Error(), ErrDBQueryFailed.Error()

	default:
		return "[[Un Caught Error]] :" + err.Error(), "[[Un Caught Error]] :" + err.Error()

	}
}
