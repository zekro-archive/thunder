package thunder_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/zekroTJA/thunder"
)

const (
	FILENAME = "test.th"
)

var db *thunder.DB

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	err := os.Remove(FILENAME)
	if !os.IsNotExist(err) && err != nil {
		panic(err)
	}
}

func TestInit(t *testing.T) {
	var err error
	db, err = thunder.Open(FILENAME)
	check(err)
}

func TestNodeCreate(t *testing.T) {
	node, err := db.CreateNode("testNode")
	check(err)
	if _, ok := db.GetNode("testNode"); !ok {
		panic("Test node could not be get after creating.")
	}
	node.Set("testKey", "testValue")
	db.Save()
	if v, ok := node.Get("testKey"); !ok {
		panic("Node value could not be get after setting.")
	} else if v != "testValue" {
		panic(fmt.Sprintf("Node value is not equal to set value: %v != testValue", v))
	}
}

func TestCheckCreatedNode(t *testing.T) {
	db = nil
	db, err := thunder.Open(FILENAME)
	check(err)
	node, ok := db.GetNode("testNode")
	if !ok {
		panic("Created node was not saved.")
	}
	value, ok := node.Get("testKey")
	if !ok {
		panic("Created value was not saved.")
	}
	if value != "testValue" {
		panic(fmt.Sprintf("Saved test value is not equal: %v != testValue", value))
	}
}

func TestRemoveNode(t *testing.T) {
	db = nil
	db, err := thunder.Open(FILENAME)
	check(err)
	err = db.RemoveNode("testNode")
	check(err)

	db = nil
	db, err = thunder.Open(FILENAME)
	check(err)
	if _, ok := db.GetNode("testNode"); ok {
		panic("Node still exists after deleting.")
	}
}
