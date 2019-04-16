package thunder

import "errors"

const (
	// MISC
	headerName    = "godb_database_file"
	headerVersion = 101
)

var (
	// ERRORS

	// ErrNodeKeyExists will be returned when a key-value pair
	// was intended to be created which key is already existent
	// in the node.
	ErrNodeKeyExists = errors.New("ndoe key already exist")
	// ErrNodeKeyNotExist will be returned when a key-value pair
	// was intended to be read by a key which is not existent.
	ErrNodeKeyNotExist = errors.New("node key does not exist")
	// ErrNodeValueNotExist will be returned when a requested key
	// exists but the key has no value.
	ErrNodeValueNotExist = errors.New("node value does not exist")
	// ErrNodeNil is returned when the node on which functions
	// are executed is nil.
	ErrNodeNil = errors.New("node is nil")
)
