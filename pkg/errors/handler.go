package errors

import (
	"log"
	"regexp"
	"strings"
)

func HandleException(err error) {
	switch err.(type) {
	case DBError:
		HandleDBException(err)

	case GcpError:
		HandleGcpException(err)

	default:
		HandleDefaultException(err)
	}
}

func HandleDefaultException(err error) {
	// check if its a sql error
	if strings.Contains(err.Error(), "sql") {
		HandleSQLException(err)
		return
	}

	// check sql state
	sqlState := getSQLState(err)
	if sqlState != "" {
		HandleSQLStateException(sqlState)
	}

	// Default exception handling
	log.Println("[[UNKNOWN ERROR]] :", err.Error())
}

func getSQLState(e error) string {
	re := regexp.MustCompile(`SQLSTATE (\w+)`)

	// Find the match in the error message
	match := re.FindStringSubmatch(e.Error())

	// Extract and return the SQLSTATE
	if len(match) > 1 {
		return match[1]
	} else {
		return ""
	}
}

