package thunder

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
	node.Data[key] = value
}
