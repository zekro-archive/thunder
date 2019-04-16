package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/zekroTJA/thunder"
)

type User struct {
	Username  string
	UID       int
	CreatedAt int64
}

func NewUser(username string, uid int) *User {
	return &User{
		Username:  username,
		UID:       uid,
		CreatedAt: time.Now().Unix(),
	}
}

func (u *User) Print() {
	fmt.Printf("Username:  %s\nUID:       %d\nCreatedAt: %d\n",
		u.Username, u.UID, u.CreatedAt)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	cinreader := bufio.NewReader(os.Stdin)

	gob.Register(new(User))

	db, err := thunder.Open("userdb.th")
	check(err)
	defer db.Save()
	defer db.Close()

	nodeUsers, ok := db.GetNode("users")
	if !ok {
		nodeUsers, err = db.CreateNode("users")
		check(err)
	}
	fmt.Println("TEST")
	defer func() {
		n, _ := db.GetNode("users")
		v, err := n.Get("test")
		fmt.Println("TEST:", v, err)
	}()

	nodeStats, ok := db.GetNode("stats")
	if !ok {
		nodeStats, err = db.CreateNode("stats")
		check(err)
	}

	fmt.Println(`
Available commands:
> get <username>
> create <username>
> delete <username>
> exit
	`)

	for {
		fmt.Print("> ")
		inpt, err := cinreader.ReadString('\n')
		check(err)
		inptsplit := strings.Split(string(inpt)[0:len(string(inpt))-1], " ")
		invoke := inptsplit[0]
		args := inptsplit[1:]

		if invoke[:len(invoke)-1] == "exit" {
			break
		}

		if len(args) < 1 {
			log.Println("[ERR] invalid number of argumnets")
		}

		switch invoke {
		case "get":
			if user, ok := nodeUsers.Get(args[0]); ok {
				user.(*User).Print()
			} else {
				log.Println("[ERR] user does not exist")
			}
		case "create":
			if _, ok := nodeUsers.Get(args[0]); ok {
				log.Println("[ERR] username taken")
			} else {
				var uid int
				if _uid, ok := nodeStats.Get("lastuid"); ok {
					uid = _uid.(int) + 1
				}
				nodeStats.Set("lastuid", uid)

				user := NewUser(args[0], uid)
				nodeUsers.Set(user.Username, user)
				check(db.Save())
				log.Println("[INFO] user created:")
				user.Print()
			}
		case "delete":
			if _, ok := nodeUsers.Get(args[0]); ok {
				err := nodeUsers.Remove(args[0])
				if err != nil {
					log.Println(err)
				}
				log.Println("user deleted")
			} else {
				log.Println("[ERR] user does not exist")
			}
		}
	}
}
