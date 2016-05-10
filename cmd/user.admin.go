package cmd

import (
	"bitbucket.org/kormaton/slapi/users/admin"
	"errors"
	"fmt"
)

type UserInvite struct {
	Args struct {
		Email string `positional-arg-name:"EMAIL" description:"User's email address"`
	} `positional-args:"yes" required:"yes"`
}

var userInvite UserInvite

func init() {
	parser.AddCommand("users.admin.invite", "Invite user to slack server", "", &userInvite)
}

func (ui *UserInvite) Execute(_ []string) error {
	ok := uadmin.Invite(ui.Args.Email)
	if !ok {
		return errors.New("users.admin.invite failure")
	}
	fmt.Println("Success.")
	return nil
}
