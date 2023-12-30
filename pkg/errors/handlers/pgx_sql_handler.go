package handlers

import (
	"error-berror/pkg/errors/models"

	"github.com/jackc/pgx/v5"
)

func handlePgxSqlException(err error) (string, string) {
	switch err {
	case pgx.ErrNoRows:
		return models.PgxSqlError{ErrMsg: err.Error()}.Error(), "No rows found"

	case pgx.ErrTooManyRows:
		return models.PgxSqlError{ErrMsg: err.Error()}.Error(), "Too many rows found"

	case pgx.ErrTxCommitRollback:
		return models.PgxSqlError{ErrMsg: err.Error()}.Error(), "Transaction commit failed"

	default:
		return "", ""
	}
}
