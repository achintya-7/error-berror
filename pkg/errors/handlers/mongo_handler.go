package handlers

import (
	"error-berror/pkg/errors/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func handleMongoException(err error) (systemErr, userErr string) {
	switch err {
	case mongo.ErrNoDocuments:
		return models.MongoError{ErrMsg: err.Error()}.Error(), "Data not found"

	case mongo.ErrNilDocument:
		return models.MongoError{ErrMsg: err.Error()}.Error(), "Nil document"

	case mongo.ErrUnacknowledgedWrite:
		return models.MongoError{ErrMsg: err.Error()}.Error(), "Unacknowledged write"

	case mongo.ErrClientDisconnected:
		return models.MongoError{ErrMsg: err.Error()}.Error(), "Mongo Client disconnected"
	
	case mongo.ErrMultipleIndexDrop:
		return models.MongoError{ErrMsg: err.Error()}.Error(), "Multiple index drop"
	
	case mongo.ErrNilValue:
		return models.MongoError{ErrMsg: err.Error()}.Error(), "Mongo Nil value"

	case mongo.ErrNilCursor:
		return models.MongoError{ErrMsg: err.Error()}.Error(), "Mongo Nil cursor"

	default:
		return "", ""
	}
}
