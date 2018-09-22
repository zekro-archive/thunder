package thunder

import "encoding/gob"

type genericMap map[interface{}]interface{}

type Node struct {
	Data genericMap
}

func NewNode() *Node {
	return &Node{Data: make(genericMap)}
}

func (node *Node) Get(key interface{}) (interface{}, bool) {
	value, ok := node.Data[key]
	return value, ok
}

func (node *Node) Set(key, value interface{}) {
	gob.Register(value)
	node.Data[key] = value
}

func (node *Node) Remove(key interface{}) error {
	if _, ok := node.Data[key]; !ok {
		return ERR_NODE_VALUE_NOT_EXISTS
	}
	delete(node.Data, key)
	return nil
}
