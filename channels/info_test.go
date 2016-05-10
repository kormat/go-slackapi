package channels

import (
	"github.com/kormat/go-slackapi/test"
	"io/ioutil"
	"testing"
)

func TestParseInfo(t *testing.T) {
	var input, err = ioutil.ReadFile("../test/channels.info.json")
	if err != nil {
		t.Errorf("Unable to read file")
	}
	var tm = test.TestMeta{t, true}
	c, ok := parseInfo(input)
	if !ok {
		t.Fatalf("Unable to parse file: %v", err)
	}
	tm.Eq("ID", "C165BUACU", c.Id)
	tm.Eq("Name", "general", c.Name)
	tm.Eq("Channel flag", true, c.IsChannel)
	tm.Eq("Created", 1462383037, c.Created)
	tm.Eq("Creator", "U165E60A2", c.Creator)
	tm.Eq("Archived flag", false, c.IsArchived)
	tm.Eq("General flag", true, c.IsGeneral)
	tm.Eq("Member flag", true, c.IsMember)
	tm.Eq("Last read", 1462387257.000026, c.LastRead)
	tm.Eq("Latest", latest{
		"message", "U165E60A2", "<@U165T1UMT>: thank you for spotting that.",
		1462387257.000026}, c.Latest)
	tm.Eq("Unread count", 42, c.UnreadCount)
	tm.Eq("Unread count display", 24, c.UnreadCountDisplay)
	tm.Eq("Members", []string{"U165E60A2", "U165N9BKJ", "U165S54BF", "U165T1UMT", "U165XNKB3"}, c.Members)
	tm.Eq("Topic", TopicPurpose{"Company-wide announcements and work-based matters", "", 0}, c.Topic)
	tm.Eq("Purpose", TopicPurpose{"This has no purpose.", "U165E60A2", 1462380000}, c.Purpose)
	if !tm.Ok {
		t.Log(c.String())
	}
}
