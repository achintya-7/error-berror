package handlers

import (
	"fmt"

	"github.com/achintya-7/error-berror/models"
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
		return models.PGXError{ErrMsg: err.Message + "Constraint at : " + err.ConstraintName}.Error(), "Column already exists"

	case pgerrcode.ForeignKeyViolation:
		return models.PGXError{ErrMsg: err.Message + "Contraint at : " + err.ConstraintName}.Error(), "Foreign key violation"

	case pgerrcode.DuplicateObject:
		return models.PGXError{ErrMsg: err.Message}.Error(), fmt.Sprintf("Duplicate object : %s", err.ConstraintName)

	case pgerrcode.UniqueViolation:
		return models.PGXError{ErrMsg: err.Message}.Error(), fmt.Sprintf("Duplicate foreign key issue with : %s", err.ConstraintName)

	default:
		return "[Unhandled PGX Error] :" + "[Code] =>" + err.Code + "[Message] =>" + err.Message + "[Query] =>" + err.InternalQuery, "Service is down, please try again later"
	}

}
