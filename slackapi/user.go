package main

import (
	"errors"
	"fmt"
	"github.com/kormat/go-slackapi/users"
)

type UserInfo struct {
	Args struct {
		ID string `description:"User ID"`
	} `positional-args:"yes" required:"yes"`
}
type UserList struct{}

var userInfo UserInfo
var userList UserList

func init() {
	parser.AddCommand("users.info", "Show user info", "", &userInfo)
	parser.AddCommand("users.list", "List users", "", &userList)
}

func (u *UserInfo) Execute(_ []string) error {
	info, ok := users.Info(u.Args.ID)
	if !ok {
		return errors.New("users.info failure")
	}
	fmt.Println(info)
	return nil
}

func (ul *UserList) Execute(_ []string) error {
	ulist, ok := users.List()
	if !ok {
		return errors.New("users.list failure")
	}
	for i, u := range ulist {
		fmt.Printf("%d. `%s` (Id: %s)\n", i, u.Name, u.Id)
	}
	return nil
}
