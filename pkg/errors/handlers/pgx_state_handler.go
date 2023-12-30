package handlers

import (
	"error-berror/pkg/errors/models"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func handlePgxStateException(err *pgconn.PgError) (string, string) {
	switch err.Code {
	case pgerrcode.ConnectionFailure:
		return models.ErrDBConnectionFailed.Error(), models.ErrDBConnectionFailed.Error()

	case pgerrcode.QueryCanceled:
		return models.PGXError{ErrMsg: err.Message}.Error(), "Query cancelled"

	case pgerrcode.SyntaxError:
		return models.PGXError{ErrMsg: err.Message}.Error(), "Syntax error in query"

	case pgerrcode.UndefinedTable:
		return models.PGXError{ErrMsg: err.Message}.Error(), "Invalid table name"

	case pgerrcode.DuplicateTable:
		return models.PGXError{ErrMsg: err.Message}.Error(), "Table already exists"

	case pgerrcode.DuplicateColumn:
		return models.PGXError{ErrMsg: err.Message}.Error(), "Column already exists"

	case pgerrcode.DuplicateObject:
		return models.PGXError{ErrMsg: err.Message}.Error(), "Object already exists"

	default:
		return "[Uncaught PGX Error] :" + "[Code] =>" + err.Code + "[Message] =>" + err.Message + "[Query] =>" + err.InternalQuery, "Service is down, please try again later"
	}

}
