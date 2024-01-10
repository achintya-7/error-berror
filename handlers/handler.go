package handlers

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func HandleDefaultException(err error) (systemErr, userErr string) {
	// check for sql state error
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return handlePgxStateException(pgErr)
	}

	// check for pgx sql error
	systemErr, userErr = handlePgxSqlException(err)
	if systemErr != "" {
		return
	}

	// check for mongo error
	systemErr, userErr = handleMongoException(err)
	if systemErr != "" {
		return
	}

	// Default exception handling
	return "[[Unhandled Error]] :" + err.Error(), "Service is down, please try again later"
}
