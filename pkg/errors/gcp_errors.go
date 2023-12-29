package errors

import (
	"errors"
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

func HandleGcpException(err error) (string, string) {
	switch err {
	case ErrGcpConnectionFailed:
		return ErrGcpConnectionFailed.Error(), ErrGcpConnectionFailed.Error()

	case ErrGcpQueryFailed:
		return ErrGcpQueryFailed.Error(), "Unable to fetch data from GCP"

	case ErrGcpPubsubFailed:
		return ErrGcpPubsubFailed.Error(), "Unable to publish message to GCP"

	default:
		return "[[Un Caught Error]] :" + err.Error(), "Service is down, please try again later"

	}
}
