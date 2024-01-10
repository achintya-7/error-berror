package exceptionsgo

import (
	"github.com/achintya-7/error-berror/handlers"
	"github.com/achintya-7/error-berror/models"
)

func HandleException(err error) (systemErr, userErr string) {
	switch err.(type) {
	case models.DBError:
		return handlers.HandleDBException(err)

	case models.GcpError:
		return handlers.HandleGcpException(err)

	default:
		return handlers.HandleDefaultException(err)
	}
}
