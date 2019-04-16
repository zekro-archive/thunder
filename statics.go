package thunder

import "errors"

const (
	// MISC
	headerName    = "godb_database_file"
	headerVersion = 101
)

var (
	// ERRORS
	ErrNodeKeyExists     = errors.New("ndoe key already exist")
	ErrNodeKeyNotExist   = errors.New("node key does not exist")
	ErrNodeValueNotExist = errors.New("node value does not exist")
	ErrNodeNil           = errors.New("node is nil")
)
