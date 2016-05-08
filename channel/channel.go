package channel

import (
	"encoding/json"
	"github.com/kr/pretty"
)

type Channel struct {
	Id                   string
	Name                 string
	IsChannel            bool `json:"is_channel"`
	Created              int
	Creator              string
	IsArchived           bool `json:"is_archived"`
	IsGeneral            bool `json:"is_general"`
	Members              []string
	Topic                topic_purpose
	Purpose              topic_purpose
	IsMember             bool `json:"is_member"`
	Last_read            string
	Latest               latest
	Unread_count         int
	Unread_count_display int
}

type topic_purpose struct {
	Value    string
	Creator  string
	Last_set int
}

type latest struct {
	Type string
	User string
	Text string
	TS   string
}

func parse(data []byte) (Channel, error) {
	var c Channel
	err := json.Unmarshal(data, &c)
	return c, err
}

func (c *Channel) String() string {
	return pretty.Sprint(c)
}
