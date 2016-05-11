package main

import (
	"fmt"
	"github.com/kormat/go-slackapi/channels"
	"github.com/kormat/go-slackapi/config"
)

type ChanInfo struct {
	Args struct {
		ID string `description:"Channel ID"`
	} `positional-args:"yes" required:"yes"`
}
type ChanList struct{}
type ChanInvite struct {
	Args struct {
		Channel string `positional-arg-name:"CHAN" description:"Channel ID"`
		User    string `positional-arg-name:"USER" description:"User ID"`
	} `positional-args:"yes" required:"yes"`
}

type ChanKick struct {
	ChanInvite
}

var chanInfo ChanInfo
var chanList ChanList
var chanInvite ChanInvite
var chanKick ChanKick

func init() {
	parser.AddCommand("channels.info", "Show channel info", "", &chanInfo)
	parser.AddCommand("channels.list", "List channels", "", &chanList)
	parser.AddCommand("channels.invite", "Invite user to channel", "", &chanInvite)
	parser.AddCommand("channels.kick", "Kick user from channel", "", &chanKick)
}

func (c *ChanInfo) Execute(_ []string) error {
	if config.CfgErr != nil {
		return config.CfgErr
	}
	info, err := channels.Info(c.Args.ID)
	if err != nil {
		return err
	}
	fmt.Println(info)
	return nil
}

func (c *ChanList) Execute(_ []string) error {
	if config.CfgErr != nil {
		return config.CfgErr
	}
	chans, err := channels.List()
	if err != nil {
		return err
	}
	for i, c := range chans {
		fmt.Printf("%d. `%s` (Id: %s)\n", i, c.Name, c.Id)
	}
	return nil
}

func (c *ChanInvite) Execute(_ []string) error {
	return channels.Invite(c.Args.Channel, c.Args.User)
}

func (c *ChanKick) Execute(_ []string) error {
	return channels.Kick(c.Args.Channel, c.Args.User)
}
