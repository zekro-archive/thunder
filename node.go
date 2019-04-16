package thunder

import (
	"encoding/gob"
	"sync"
)

type genericMap map[interface{}]interface{}

// A Node contains the actual key-value data
// map where data handled with.
type Node struct {
	mx     *sync.Mutex
	locked bool

	Data genericMap
}

func (node *Node) lock() {
	if node.mx != nil && !node.locked {
		node.mx.Lock()
		node.locked = true
	}
}

func (node *Node) unlock() {
	if node.mx != nil && node.locked {
		node.locked = false
		node.mx.Unlock()
	}
}

// NewNode initializes a new node.
func NewNode() *Node {
	return &Node{
		Data: make(genericMap),
	}
}

// Get returns the value of the node by key.
// If there is no value existent, nil and false
// is returned.
func (node *Node) Get(key interface{}) (interface{}, bool) {
	if node == nil {
		return nil, false
	}

	value, ok := node.Data[key]
	return value, ok
}

// Set sets the passed value to the passed key.
// If the node pointer is nil, this returns an
// ErrNodeNil error.
func (node *Node) Set(key, value interface{}) error {
	if node == nil {
		return ErrNodeNil
	}

	gob.Register(value)
	node.Data[key] = value

	return nil
}

// Remove deletes a key-value pair by the
// key. If the node pointer is nil, this
// returns an ErrNodeNil error.
func (node *Node) Remove(key interface{}) error {
	if node == nil {
		return ErrNodeNil
	}

	if _, ok := node.Data[key]; !ok {
		return ErrNodeValueNotExist
	}

	delete(node.Data, key)
	return nil
}
