package thunder_test

import (
	"fmt"
	"os"
	"testing"

	thunder ".."
)

func init() {
	err := os.Remove(FILENAME)
	if !os.IsNotExist(err) && err != nil {
		panic(err)
	}
}

func TestComplex(t *testing.T) {
	var err error
	db, err = thunder.Open(FILENAME)
	check(err)

	node, err := db.CreateNode(1)
	check(err)

	testMap := make(map[int]string)
	testMap[1] = "test1231231"

	node.Set("test1", testMap)
	node.Set("test2", 123)

	db.Close()

	db, err = thunder.Open(FILENAME)
	check(err)

	node, _ = db.GetNode(1)

	fmt.Println(node)
}
