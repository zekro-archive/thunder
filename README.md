---

# Introduction

`thunder` is a small and lightweight database and storage package for Go. This package is constipated for small data saving for your application, simply structured by nodes containing key-value pair tables, without keeping formatting and serializing in mind. Just save and load your data as you use it in your application. **Attention:** You should know, that the full database file will be loaded into the system memory on activation. So the memory usage will scale with the size of the database. This package is not designed for the huge database!

---

# Usage

Get the package with
```
go get github.com/zekroTJA/thunder
```

Then, create a new database by opening a non-existing file. This will always automatically create a new and empty database:

```go

package main

import (
    "log"

    "github.com/zekroTJA/thunder"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    db, err := thunder.Open("myDatabase.th")
    check(err)
    defer db.Close()

    node, err := db.CreateNode("users")
    check(err)

    username := "zekro"
    uid := 45687236137813

    // Every time you set a value in a note, the
    // value will aslo be set in the database variable.
    // To ensure that all data is saved to file, also if
    // the application crashes and 'db.Close()' will not
    // be executed, execute 'db.Save()' every time after
    // setting values in the nodes. Node creation and
    // deletion will be automatically saved to database.
    node.Set(username, uid)
    check(db.Save())

}

```

If you are using custom structs or *complex'* types, you need to register them in `gob` that the database can be deserialized properly.

```go

package main

import (
    "log"
    "encoding/gob"

    "github.com/zekroTJA/thunder"
)

type User struct {
    Username string
    UID      int64
}

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    // register type in gob for later deserialization.
    gob.Register(map[int64]*User{})

    db, err := thunder.Open("mySecondDatabase.th")
    check(err)
    defer db.Close()

    registeredUsers := make(map[int64]*User)
    registeredUsers[5434524234234234] = &User{
        Username: "zekro",
        UID:      5434524234234234,
    }

    node, ok := GetNode("users")
    if !ok {
        node, err := db.CreateNode("users")
        check(err)
    }
    node.Set("registeredUsers", registeredUsers)
    check(db.Save())
}
```

More examples are contained inside the `examples` folder in this repository.

---

© 2018 zekro Development  

[zekro.de](https://zekro.de) | contact[at]zekro.de
