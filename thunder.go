package thunder

import (
	"encoding/gob"
	"errors"
	"io"
	"os"
	"sync"
)

type header struct {
	Name    string
	Version int
}

type nodeMap map[interface{}]*Node

// DB is the struct that contains the name of the of the database file,
// the header containing database type and version and the
// data as map of data nodes.
type DB struct {
	mx     *sync.Mutex
	locked bool

	filename string
	header   *header
	data     nodeMap
}

// ------------------------ PRIVATE FUNCS ------------------------

// encode transmits the data of the DB instance
// to the passed file writer handler.
func encode(db *DB, fhandler io.Writer) error {
	gobencoder := gob.NewEncoder(fhandler)
	err := gobencoder.Encode(db)
	return err
}

// decode reads the data from the passed file
// reader handler and parses th data to a new
// instance of DB.
func decode(fhandler io.Reader) (*DB, error) {
	gobdecoder := gob.NewDecoder(fhandler)
	obj := new(DB)
	err := gobdecoder.Decode(obj)
	return obj, err
}

// lock locks the mutex of the DB if it is
// not already locked.
func (db *DB) lock() {
	if db.mx != nil && !db.locked {
		db.mx.Lock()
		db.locked = true
	}
}

// unlock unlocks the mutex of the DB if
// it is locked.
func (db *DB) unlock() {
	if db.mx != nil && db.locked {
		db.locked = false
		db.mx.Unlock()
	}
}

// ------------------------ PUBLIC FUNCS ------------------------

// Register records one or more types passed by an
// instance of them to gob.
// See gob#Register for more information.
func Register(val ...interface{}) {
	for _, v := range val {
		gob.Register(v)
	}
}

// Open creates a new instance of DB from database file.
// If the passed file does not exist, it will be created
// as empty database.
// The file name and location is passed as string.
// If no exceptions are occuring, the database instannce
// will be returned. Else, the error will be returned as
// second return value.
//
// If you get an error like:
// "gob: name not registered for interface: <type>"
// You need to register this type in gob before opening
// like following:
//
//	import (
//		"encoding/gob"
//		"github.com/zekroTJA/thunder"
//	)
//
//	type User struct {
// 		UserName string
//		UID      int64
// 	}
//
// 	func main() {
//		thunder.Register(*User)
//		db, err := thunder.Open("myDb.th")
//	}
func Open(filename string) (*DB, error) {
	gob.Register(map[interface{}]*Node{})

	fhandler, err := os.Open(filename)
	if os.IsNotExist(err) {
		obj := &DB{
			mx: new(sync.Mutex),
			header: &header{
				Name:    headerName,
				Version: headerVersion,
			},
			data:     make(nodeMap),
			filename: filename,
		}
		fhandler, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
		defer fhandler.Close()
		err = encode(obj, fhandler)
		return obj, err
	} else if err != nil {
		return nil, err
	}
	defer fhandler.Close()

	obj, err := decode(fhandler)
	if err != nil {
		return nil, err
	}
	if obj.header.Version > headerVersion {
		return nil, errors.New("the database file version is newer than the package version")
	}

	obj.filename = filename
	return obj, err
}

// CreateNode creates a new node inside data base.
// Notes are used to save key-value pair data.
// As first arguement, the key will be passed as interface type.
// Optional, the externally created node can be passed to
// insert prepared nodes into the data base.
// If no exceptions occure, the created node instance will be returned.
// Else, the error will be returned as second return value.
func (db *DB) CreateNode(key interface{}, node ...*Node) (*Node, error) {
	db.lock()
	defer db.unlock()

	if _, ok := db.data[key]; ok {
		return nil, ErrNodeKeyExists
	}
	if len(node) > 0 {
		db.data[key] = node[0]
	} else {
		db.data[key] = NewNode()
	}
	db.Save()
	return db.data[key], nil
}

// GetNode gets the node by key if exists.
// It returns the node instance of the key and,
// as bool, if the node key exists in the database.
func (db *DB) GetNode(key interface{}) (*Node, bool) {
	db.lock()
	defer db.unlock()

	node, ok := db.data[key]
	return node, ok
}

// RemoveNode deletes the node by key in the database.
// If erros occure, they will be returned as error.
func (db *DB) RemoveNode(key interface{}) error {
	db.lock()
	defer db.unlock()

	if _, ok := db.data[key]; !ok {
		return ErrNodeKeyNotExist
	}
	delete(db.data, key)
	db.Save()
	return nil
}

// Save saves the current database state to file.
// If errors occure, they will be returned as error.
func (db *DB) Save() error {
	db.lock()
	defer db.unlock()

	fhandler, err := os.OpenFile(db.filename, os.O_WRONLY, 771)
	defer fhandler.Close()
	if err != nil {
		return err
	}
	err = encode(db, fhandler)
	return err
}

// Close closes the database file and saves the current
// state of the database to the file.
func (db *DB) Close() {
	db.Save()
}
