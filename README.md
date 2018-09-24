<div align="center">
     <!-- <img src="https://zekro.de/src/go_chat_logo.png" width="400"/> -->
     <h1>~ thunder ~</h1>
     <strong>Small and fast database package for Go</strong><br><br>
     <img src="https://forthebadge.com/images/badges/made-with-go.svg" height="30" />&nbsp;
     <!-- <img src="https://forthebadge.com/images/badges/60-percent-of-the-time-works-every-time.svg" height="30" />&nbsp; -->
     <a href="https://travis-ci.org/zekroTJA/thunder"><img src="https://img.shields.io/travis/zekroTJA/thunder.svg?style=for-the-badge&logo=travis" height="30"></a>&nbsp;
     <a href="https://godoc.org/github.com/zekroTJA/thunder"><img src="https://img.shields.io/badge/docs-godoc-0ee6ea.svg?style=for-the-badge" height="30"></a>&nbsp;
     <a href="https://zekro.de/discord"><img src="https://img.shields.io/discord/307084334198816769.svg?logo=discord&style=for-the-badge" height="30"></a>
</div>

---

# Introduction

`thunder` is a small, lightweight database and storage package for Go. This package is concipated for small data saving for your applications, simply structured by nodes containing key-value pair tables, without mind of formatting or serializing. Just save and load your data like you use it in your application. **Attention:** You sould know, that the full database file will be loaded in the memory on usage. So the memory usage will scale with the size of the database. This package is not designed for huge database scales!

---

# Usage

Get the package with
```
go get github.com/zekroTJA/thunder
```

Then, cretae a new database by opening a non existing file. This will always automaticall create a new and empty database:

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

If you are using custom structs or *'complex'* types, you need to register them in `gob` that the database can be deserialized properly.

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

You can find some more examples in the `examples` folder in this repository.

---

© 2018 zekro Development  

[zekro.de](https://zekro.de) | contact[at]zekro.de


