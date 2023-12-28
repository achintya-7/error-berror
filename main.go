package main

import (
	"database/sql"
	err "error-berror/pkg/errors"
	"errors"
)

func main() {
	err.HandleDBException(err.ErrDBQueryFailed)
	err.HandleException(sql.ErrNoRows)
	err.HandleException(errors.New("some error"))

	println()
}
