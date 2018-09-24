package thunder_test

import (
	"fmt"
	"os"
	"testing"

	thunder ".."
)

type User struct {
	Username string
	Uid      int64
}

func init() {
	err := os.Remove(FILENAME)
	if !os.IsNotExist(err) && err != nil {
		panic(err)
	}
}

func TestComplex(t *testing.T) {
	var err error
	var node *thunder.Node

	thunder.Register(map[int]*User{})

	db, err = thunder.Open(FILENAME)
	check(err)

	node, err = db.CreateNode("t")
	check(err)

	testMap := map[int]*User{}
	testMap[1] = &User{"Herbert", 1234567890}

	node.Set("test1", testMap)
	node.Set("test2", 123)

	db.Close()

	db, err = thunder.Open(FILENAME)
	check(err)

	node, ok := db.GetNode("t")

	fmt.Println(node, ok)
}
