package errors

import (
	"errors"
	"log"
)

type GcpError struct {
	ErrMsg string
}

// Custom GCP errors
var (
	ErrGcpConnectionFailed = GcpError{ErrMsg: "Gcp connection failed"}
	ErrGcpQueryFailed      = GcpError{ErrMsg: "Gcp query failed"}
	ErrGcpPubsubFailed     = GcpError{ErrMsg: "Pubsub request failed"}
)

func (e GcpError) Error() string {
	return "[DB ERROR] : " + e.ErrMsg
}

func (e GcpError) UnWrap() error {
	return errors.New(e.ErrMsg)
}

func HandleGcpException(err error) {
	switch err {
	case ErrGcpConnectionFailed:
		log.Println(ErrGcpConnectionFailed.ErrMsg)

	case ErrGcpQueryFailed:
		log.Println(ErrGcpQueryFailed.ErrMsg)

	default:
		log.Println("Un Caught Error : ", err.Error())

	}
}