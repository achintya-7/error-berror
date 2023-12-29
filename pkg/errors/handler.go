package errors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func HandleException(err error) (systemErr, userErr string) {
	switch err.(type) {
	case DBError:
		return HandleDBException(err)

	case GcpError:
		return HandleGcpException(err)

	default:
		return HandleDefaultException(err)
	}
}

func HandleDefaultException(err error) (systemErr, userErr string) {
	// check for sql state error
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return HandlePGXException(pgErr)
	}

	// Default exception handling
	return "[[Un Caught Error]] :" + err.Error(), "Service is down, please try again later"
}
