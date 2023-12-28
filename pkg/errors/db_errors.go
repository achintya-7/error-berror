package errors

import (
	"errors"
	"log"
)

type DBError struct {
	ErrMsg string
}

// Custom DB errors
var (
	ErrDBConnectionFailed = DBError{ErrMsg: "DB connection failed"}
	ErrDBQueryFailed = DBError{ErrMsg: "DB query failed"}
	ErrRowNotFound = DBError{ErrMsg: "Not found"}
	ErrRowDuplicate = DBError{ErrMsg: "Duplicate row"}
)

func (e DBError) Error() string {
	return "[DB ERROR] : " + e.ErrMsg
}

func (e DBError) UnWrap() error {
	return errors.New(e.ErrMsg)
}

func HandleDBException(err error) {
	switch err {
	case ErrDBConnectionFailed:
		log.Println(ErrDBConnectionFailed.Error())
	
	case ErrDBQueryFailed:
		log.Println(ErrDBQueryFailed.Error())

	case ErrRowNotFound:
		log.Println(ErrRowNotFound.Error())
	
	case ErrRowDuplicate:
		log.Println(ErrRowDuplicate.Error())
	
	default:
		log.Println("Un Caught Error : ", err.Error())
		
	}
}