package handlers

import (
	"error-berror/pkg/errors/models"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func HandleException(err error) (systemErr, userErr string) {
	switch err.(type) {
	case models.DBError:
		return handleDBException(err)

	case models.GcpError:
		return handleGcpException(err)

	default:
		return handleDefaultException(err)
	}
}

func handleDefaultException(err error) (systemErr, userErr string) {
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
	return "[[Un Caught Error]] :" + err.Error(), "Service is down, please try again later"
}
