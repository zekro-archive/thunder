package thunder

import "errors"

const (
	// headerName indicates thunder that the read
	// gob file is realy a thunder database file.
	headerName = "godb_database_file"
	// headerversion is to check if the version of
	// thunder the database was created with a
	// version which si not backwards compatible.
	// So this version is not th actual thunder
	// version because it is only increased when
	// the database structure has changed.
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
