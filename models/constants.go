package models


// Custom DB errors
var (
	ErrDBConnectionFailed = DBError{ErrMsg: "DB connection failed"}
	ErrDBQueryFailed      = DBError{ErrMsg: "DB query failed"}
)

// Custom GCP errors
var (
	ErrGcpConnectionFailed   = GcpError{ErrMsg: "Gcp connection failed"}
	ErrGcpQueryFailed        = GcpError{ErrMsg: "Gcp query failed"}
	ErrGcpPubsubFailed       = GcpError{ErrMsg: "Pubsub request failed"}
	ErrSizeCalculationFailed = GcpError{ErrMsg: "Size calculation failed for file in storage"}
	ErrSignedUrlFailed       = GcpError{ErrMsg: "Signed url generation failed"}
	ErrAcknowledgeFailed     = GcpError{ErrMsg: "CRC acknowledgement failed"}
)
