package handlers

import (
	"github.com/achintya-7/error-berror/models"
)

func HandleDBException(err error) (string, string) {
	switch err {
	case models.ErrDBConnectionFailed:
		return models.ErrDBConnectionFailed.Error(), models.ErrDBConnectionFailed.Error()

	case models.ErrDBQueryFailed:
		return models.ErrDBQueryFailed.Error(), models.ErrDBQueryFailed.Error()

	default:
		return "[[Unhandled Error]] :" + err.Error(), "[[Unhandled Error]] :" + err.Error()

	}
}
