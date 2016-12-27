package main

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

type User struct {
	Name string
}

type Context struct {
	User *User
}

func (c *Context) UsersList(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, "users")
}

func (c *Context) Root(rw web.ResponseWriter, req *web.Request) {
	if c.User != nil {
		fmt.Fprint(rw, "Hello ", c.User.Name)
	} else {
		fmt.Fprint(rw, "Hello guest")
	}
}

func main() {
	router := web.New(Context{})
	router.Get("/users", (*Context).UsersList)
	router.Get("/", (*Context).Root)
	http.ListenAndServe("localhost:3000", router)
}
