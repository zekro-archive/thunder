package thunder

import (
	"encoding/gob"
	"errors"
	"io"
	"os"
)

type header struct {
	Name    string
	Version int
}

type nodeMap map[interface{}]*Node

// DB is the struct containing the file name of the database,
// the header containing database type and version and the
// data as map of data nodes.
type DB struct {
	Filename string
	Header   *header
	Data     nodeMap
}

// ------------------------ PRIVATE FUNCS ------------------------

func encode(db *DB, fhandler io.Writer) error {
	gobencoder := gob.NewEncoder(fhandler)
	err := gobencoder.Encode(db)
	return err
}

func decode(fhandler io.Reader) (*DB, error) {
	gobdecoder := gob.NewDecoder(fhandler)
	obj := &DB{}
	err := gobdecoder.Decode(obj)
	return obj, err
}

// ------------------------ PUBLIC FUNCS ------------------------

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
//		import (
//			"encoding/gob"
//			"github.com/zekroTJA/thunder"
//		)
//		type User struct {
// 			username string
//			uid      int64
// 		}
//
//  	func main() {
//			gob.Register(map[string]*User)
//			db, err := thunder.Open("myDb.th")
//		}
func Open(filename string) (*DB, error) {
	gob.Register(map[interface{}]*Node{})

	fhandler, err := os.Open(filename)
	if os.IsNotExist(err) {
		obj := &DB{
			Header: &header{
				Name:    HEADER_NAME,
				Version: HEADER_VERSION,
			},
			Data:     make(nodeMap),
			Filename: filename,
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
	if obj.Header.Version > HEADER_VERSION {
		return nil, errors.New("The database file version is newer than the package version. Please update your package to read the database file.")
	}
	obj.Filename = filename
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
	if _, ok := db.Data[key]; ok {
		return nil, ERR_NODE_KEY_EXISTS
	}
	if len(node) > 0 {
		db.Data[key] = node[0]
	} else {
		db.Data[key] = NewNode()
	}
	db.Save()
	return db.Data[key], nil
}

// GetNode gets the node by key if exists.
// It returns the node instance of the key and,
// as bool, if the node key exists in the database.
func (db *DB) GetNode(key interface{}) (*Node, bool) {
	node, ok := db.Data[key]
	return node, ok
}

// RemoveNode deletes the node by key in the database.
// If erros occure, they will be returned as error.
func (db *DB) RemoveNode(key interface{}) error {
	if _, ok := db.Data[key]; !ok {
		return ERR_NODE_KEY_NOT_EXISTS
	}
	delete(db.Data, key)
	db.Save()
	return nil
}

// Save saves the current database state to file.
// If errors occure, they will be returned as error.
func (db *DB) Save() error {
	fhandler, err := os.OpenFile(db.Filename, os.O_WRONLY, 771)
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
