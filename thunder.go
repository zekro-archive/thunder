package thunder

import (
	"encoding/gob"
	"io"
	"os"
)

type header struct {
	Name    string
	Version int
}

type nodeMap map[interface{}]*Node

type DB struct {
	Filename string
	Header   *header
	Data     nodeMap
}

func Open(filename string) (*DB, error) {
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
	obj.Filename = filename
	return obj, err
}

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

func (db *DB) GetNode(key interface{}) (*Node, bool) {
	node, ok := db.Data[key]
	return node, ok
}

func (db *DB) RemoveNode(key interface{}) error {
	if _, ok := db.Data[key]; !ok {
		return ERR_NODE_KEY_NOT_EXISTS
	}
	delete(db.Data, key)
	db.Save()
	return nil
}

func (db *DB) Save() error {
	fhandler, err := os.OpenFile(db.Filename, os.O_WRONLY, 771)
	defer fhandler.Close()
	if err != nil {
		return err
	}
	err = encode(db, fhandler)
	return err
}

func (db *DB) Close() {
	db.Save()
}

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
