package handlers

import (
	ex "error-berror/pkg/errors/models"
)

func handleGcpException(err error) (systemErr, userErr string) {
	switch err {
	case ex.ErrGcpConnectionFailed:
		return ex.ErrGcpConnectionFailed.Error(), ex.ErrGcpConnectionFailed.Error()

	case ex.ErrGcpQueryFailed:
		return ex.ErrGcpQueryFailed.Error(), "Unable to fetch data from GCP"

	case ex.ErrGcpPubsubFailed:
		return ex.ErrGcpPubsubFailed.Error(), "Unable to publish message to GCP"

	case ex.ErrSizeCalculationFailed:
		return ex.ErrSizeCalculationFailed.Error(), "Cloud storage size calculation failed"

	case ex.ErrSignedUrlFailed:
		return ex.ErrSignedUrlFailed.Error(), "Unable to create signed url for file upload"

	case ex.ErrAcknowledgeFailed:
		return ex.ErrAcknowledgeFailed.Error(), "Uploaded file got corrupted, please try again"

	default:
		return "[[Un Caught Error]] :" + err.Error(), "Service is down, please try again later"

	}
}
