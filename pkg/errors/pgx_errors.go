package errors

import (
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

type PGXError struct {
	ErrMsg string
}

func (e PGXError) Error() string {
	return "[PGX ERROR] : " + e.ErrMsg
}

func HandlePGXException(err *pgconn.PgError) (string, string) {
	switch err.Code {
	case pgerrcode.ConnectionFailure:
		return ErrDBConnectionFailed.Error(), ErrDBConnectionFailed.Error()

	case pgerrcode.QueryCanceled:
		return PGXError{ErrMsg: err.Message}.Error(), "Query cancelled"

	case pgerrcode.SyntaxError:
		return PGXError{ErrMsg: err.Message}.Error(), "Syntax error in query"

	case pgerrcode.UndefinedTable:
		return PGXError{ErrMsg: err.Message}.Error(), "Invalid table name"

	case pgerrcode.DuplicateTable:
		return PGXError{ErrMsg: err.Message}.Error(), "Table already exists"

	case pgerrcode.DuplicateColumn:
		return PGXError{ErrMsg: err.Message}.Error(), "Column already exists"

	case pgerrcode.DuplicateObject:
		return PGXError{ErrMsg: err.Message}.Error(), "Object already exists"

	default:
		return "[Un PGX Error] :" + "[Code] =>" + err.Code + "[Message] =>" + err.Message + "[Query] =>" + err.InternalQuery, "Service is down, please try again later"
	}

}
