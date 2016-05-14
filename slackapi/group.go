package main

import (
	"fmt"
	"github.com/kormat/go-slackapi/config"
	"github.com/kormat/go-slackapi/groups"
)

type GroupInfo struct {
	Args struct {
		ID string `description:"Group ID"`
	} `positional-args:"yes" required:"yes"`
}
type GroupList struct{}
type GroupInvite struct {
	Args struct {
		Group string `positional-arg-name:"CHAN" description:"Group ID"`
		User  string `positional-arg-name:"USER" description:"User ID"`
	} `positional-args:"yes" required:"yes"`
}

type GroupKick struct {
	GroupInvite
}

var groupInfo GroupInfo
var groupList GroupList
var groupInvite GroupInvite
var groupKick GroupKick

func init() {
	parser.AddCommand("groups.info", "Show group info", "", &groupInfo)
	parser.AddCommand("groups.list", "List groups", "", &groupList)
	parser.AddCommand("groups.invite", "Invite user to group", "", &groupInvite)
	parser.AddCommand("groups.kick", "Kick user from group", "", &groupKick)
}

func (g *GroupInfo) Execute(_ []string) error {
	if config.CfgErr != nil {
		return config.CfgErr
	}
	info, err := groups.Info(g.Args.ID)
	if err != nil {
		return err
	}
	fmt.Println(info)
	return nil
}

func (gl *GroupList) Execute(_ []string) error {
	if config.CfgErr != nil {
		return config.CfgErr
	}
	glist, err := groups.List()
	if err != nil {
		return err
	}
	for i, g := range glist {
		fmt.Printf("%d. `%s` (Id: %s)\n", i, g.Name, g.ID)
	}
	return nil
}

func (g *GroupInvite) Execute(_ []string) error {
	if config.CfgErr != nil {
		return config.CfgErr
	}
	return groups.Invite(g.Args.Group, g.Args.User)
}

func (g *GroupKick) Execute(_ []string) error {
	if config.CfgErr != nil {
		return config.CfgErr
	}
	return groups.Kick(g.Args.Group, g.Args.User)
}
