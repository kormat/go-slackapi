package main

import (
	"fmt"
	"github.com/kormat/go-slackapi/config"
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
	if config.CfgErr != nil {
		return config.CfgErr
	}
	info, err := users.Info(u.Args.ID)
	if err != nil {
		return err
	}
	fmt.Println(info)
	return nil
}

func (ul *UserList) Execute(_ []string) error {
	if config.CfgErr != nil {
		return config.CfgErr
	}
	ulist, err := users.List()
	if err != nil {
		return err
	}
	for i, u := range ulist {
		fmt.Printf("%d. `%s` (Id: %s)\n", i, u.Name, u.Id)
	}
	return nil
}
