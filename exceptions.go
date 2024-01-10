package exceptionsgo

import (
	"errors"
	"fmt"

	"github.com/achintya-7/error-berror/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/mongo"
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
	return "[[Unhandled Error]] :" + err.Error(), "Service is down, please try again later"
}

func handleDBException(err error) (string, string) {
	switch err {
	case models.ErrDBConnectionFailed:
		return models.ErrDBConnectionFailed.Error(), models.ErrDBConnectionFailed.Error()

	case models.ErrDBQueryFailed:
		return models.ErrDBQueryFailed.Error(), models.ErrDBQueryFailed.Error()

	default:
		return "[[Unhandled Error]] :" + err.Error(), "[[Unhandled Error]] :" + err.Error()

	}
}

func handleGcpException(err error) (systemErr, userErr string) {
	switch err {
	case models.ErrGcpConnectionFailed:
		return models.ErrGcpConnectionFailed.Error(), models.ErrGcpConnectionFailed.Error()

	case models.ErrGcpQueryFailed:
		return models.ErrGcpQueryFailed.Error(), "Unable to fetch data from GCP"

	case models.ErrGcpPubsubFailed:
		return models.ErrGcpPubsubFailed.Error(), "Unable to publish message to GCP"

	case models.ErrSizeCalculationFailed:
		return models.ErrSizeCalculationFailed.Error(), "Cloud storage size calculation failed"

	case models.ErrSignedUrlFailed:
		return models.ErrSignedUrlFailed.Error(), "Unable to create signed url for file upload"

	case models.ErrAcknowledgeFailed:
		return models.ErrAcknowledgeFailed.Error(), "Uploaded file got corrupted, please try again"

	default:
		return "[[Unhandled Error]] :" + err.Error(), "Service is down, please try again later"

	}
}

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

func handlePgxSqlException(err error) (string, string) {
	switch err {
	case pgx.ErrNoRows:
		return models.PgxSqlError{ErrMsg: err.Error()}.Error(), "No rows found"

	case pgx.ErrTooManyRows:
		return models.PgxSqlError{ErrMsg: err.Error()}.Error(), "Too many rows found"

	case pgx.ErrTxCommitRollback:
		return models.PgxSqlError{ErrMsg: err.Error()}.Error(), "Transaction commit failed"

	default:
		return "", ""
	}
}

func handlePgxStateException(err *pgconn.PgError) (string, string) {
	switch err.Code {
	case pgerrcode.ConnectionFailure:
		return models.ErrDBConnectionFailed.Error(), models.ErrDBConnectionFailed.Error()

	case pgerrcode.QueryCanceled:
		return models.PGXError{ErrMsg: err.Message}.Error(), "Query cancelled"

	case pgerrcode.SyntaxError:
		return models.PGXError{ErrMsg: err.Message}.Error(), "Syntax error in query"

	case pgerrcode.UndefinedTable:
		return models.PGXError{ErrMsg: err.Message}.Error(), "Invalid table name"

	case pgerrcode.DuplicateTable:
		return models.PGXError{ErrMsg: err.Message}.Error(), "Table already exists"

	case pgerrcode.DuplicateColumn:
		return models.PGXError{ErrMsg: err.Message + "Constraint at : " + err.ConstraintName}.Error(), "Column already exists"

	case pgerrcode.ForeignKeyViolation:
		return models.PGXError{ErrMsg: err.Message + "Contraint at : " + err.ConstraintName}.Error(), "Foreign key violation"

	case pgerrcode.DuplicateObject:
		return models.PGXError{ErrMsg: err.Message}.Error(), fmt.Sprintf("Duplicate object : %s", err.ConstraintName)

	case pgerrcode.UniqueViolation:
		return models.PGXError{ErrMsg: err.Message}.Error(), fmt.Sprintf("Duplicate foreign key issue with : %s", err.ConstraintName)

	default:
		return "[Unhandled PGX Error] :" + "[Code] =>" + err.Code + "[Message] =>" + err.Message + "[Query] =>" + err.InternalQuery, "Service is down, please try again later"
	}

}
