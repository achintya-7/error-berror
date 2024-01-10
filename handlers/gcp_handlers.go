package handlers

import "github.com/achintya-7/error-berror/models"

func HandleGcpException(err error) (systemErr, userErr string) {
	switch err {
	case models.ErrGcpConnectionFailed:
		return models.ErrGcpConnectionFailed.Error(), models.ErrGcpConnectionFailed.Error()

	case models.ErrGcpQueryFailed:
		return models.ErrGcpQueryFailed.Error(), "Unable to fetch data from GCP"

	case models.ErrGcpPubsubFailed:
		return models.ErrGcpPubsubFailed.Error(), "Unable to publish message to GCP"

	case models.ErrSizeCalculationFailed:
		return models.ErrSizeCalculationFailed.Error(), "Cloud storage size calculation failed"

	case models.ErrSignedUrlFailed:
		return models.ErrSignedUrlFailed.Error(), "Unable to create signed url for file upload"

	case models.ErrAcknowledgeFailed:
		return models.ErrAcknowledgeFailed.Error(), "Uploaded file got corrupted, please try again"

	default:
		return "[[Unhandled Error]] :" + err.Error(), "Service is down, please try again later"

	}
}
