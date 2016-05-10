package cmd

import (
	"bitbucket.org/kormaton/slapi/groups"
	"errors"
	"fmt"
)

type GroupInfo struct {
	Args struct {
		ID string `description:"Group ID"`
	} `positional-args:"yes" required:"yes"`
}
type GroupList struct{}

var groupInfo GroupInfo
var groupList GroupList

func init() {
	parser.AddCommand("groups.info", "Show group info", "", &groupInfo)
	parser.AddCommand("groups.list", "List groups", "", &groupList)
}

func (g *GroupInfo) Execute(_ []string) error {
	info, ok := groups.Info(g.Args.ID)
	if !ok {
		return errors.New("groups.info failure")
	}
	fmt.Println(info)
	return nil
}

func (gl *GroupList) Execute(_ []string) error {
	glist, ok := groups.List()
	if !ok {
		return errors.New("groups.list failure")
	}
	for i, g := range glist {
		fmt.Printf("%d. `%s` (Id: %s)\n", i, g.Name, g.Id)
	}
	return nil
}
