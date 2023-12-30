package handlers

import (
	ex "error-berror/pkg/errors/models"
)

func handleDBException(err error) (string, string) {
	switch err {
	case ex.ErrDBConnectionFailed:
		return ex.ErrDBConnectionFailed.Error(), ex.ErrDBConnectionFailed.Error()

	case ex.ErrDBQueryFailed:
		return ex.ErrDBQueryFailed.Error(), ex.ErrDBQueryFailed.Error()

	default:
		return "[[Un Caught Error]] :" + err.Error(), "[[Un Caught Error]] :" + err.Error()

	}
}
