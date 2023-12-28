package errors

import "log"

type SqlStateError struct {
	ErrMsg string
}

func (e SqlStateError) Error() string {
	return "[SQL STATE ERROR] : " + e.ErrMsg
}

func HandleSQLStateException(err string) {
	switch err {
	case "23000":
		log.Println(SqlStateError{ErrMsg: "SQLSTATE 23000 : Integrity Constraint Violation"}.Error())

	case "42000":
		log.Println(SqlStateError{ErrMsg: "SQLSTATE 42000 : Syntax Error or Access Rule Violation"}.Error())

	default:
		log.Println(SqlStateError{ErrMsg: "Un Caught Error :" + err}.Error())

	}
}
