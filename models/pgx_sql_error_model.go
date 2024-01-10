package models

import "errors"

type PgxSqlError struct {
	ErrMsg string
}

func (e PgxSqlError) Error() string {
	return "[PGX SQL ERROR] : " + e.ErrMsg
}

func (e PgxSqlError) UnWrap() error {
	return errors.New(e.ErrMsg)
}
