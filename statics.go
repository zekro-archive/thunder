package thunder

import "errors"

const (
	// MISC
	HEADER_NAME    = "godb_database_file"
	HEADER_VERSION = 10
)

var (
	// ERRORS
	ERR_NODE_KEY_EXISTS       = errors.New("NODE_KEY_EXISTS")
	ERR_NODE_KEY_NOT_EXISTS   = errors.New("NODE_KEY_NOT_EXISTS")
	ERR_NODE_VALUE_NOT_EXISTS = errors.New("NODE_VALUE_NOT_EXISTS")
)
