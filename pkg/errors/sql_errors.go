package errors

import (
	"database/sql"
	"log"
)

type SqlError struct {
	ErrMsg string
}

func (e SqlError) Error() string {
	return "[SQL ERROR] : " + e.ErrMsg
}

func HandleSQLException(err error) {
	switch err {
	case sql.ErrNoRows:
		log.Println(SqlError{ErrMsg: sql.ErrNoRows.Error()}.Error())

	case sql.ErrConnDone:
		log.Println(SqlError{ErrMsg: sql.ErrConnDone.Error()}.Error())

	case sql.ErrTxDone:
		log.Println(SqlError{ErrMsg: sql.ErrTxDone.Error()}.Error())

	default:
		log.Println(SqlError{ErrMsg: "Un Caught Error :" + err.Error()}.Error())

	}

}
