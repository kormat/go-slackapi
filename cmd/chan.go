package cmd

import (
	"bitbucket.org/kormaton/slapi/channels"
	"errors"
	"fmt"
)

type ChanList struct{}

type ChanInfo struct {
	Args struct {
		ID string `description:"Channel ID"`
	} `positional-args:"yes" required:"yes" value-name:"uh?"`
}

var chanList ChanList
var chanInfo ChanInfo

func init() {
	parser.AddCommand("channels.list", "List channels", "", &chanList)
	parser.AddCommand("channels.info", "Show channel info", "", &chanInfo)
}

func (c *ChanList) Execute(_ []string) error {
	chans, ok := channels.List()
	if !ok {
		return errors.New("Failed to get channel list")
	}
	for i, c := range chans {
		fmt.Printf("%d. `%s` (Id: %s)\n", i, c.Name, c.Id)
	}
	return nil
}

func (c *ChanInfo) Execute(_ []string) error {
	info, ok := channels.Info(c.Args.ID)
	if !ok {
		return errors.New("channels.info failure")
	}
	fmt.Println(info)
	return nil
}
