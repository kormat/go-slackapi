package query

import (
	"github.com/kormat/go-slackapi/config"
	"net/url"
	"testing"
)

func TestCreateUrl(t *testing.T) {
	old := config.Cfg
	defer func() { config.Cfg = old }()
	config.Cfg = &config.Config{Token: "tokeny", APIEndpoint: "http://api.endpoint"}
	cases := []struct {
		in_cmd string
		in_v   url.Values
		want   string
	}{
		{"cmd0", url.Values{}, "http://api.endpoint/cmd0?token=tokeny"},
		{"cmd1", url.Values{"p1": []string{"v1"}, "p2": []string{"v2"}},
			"http://api.endpoint/cmd1?p1=v1&p2=v2&token=tokeny"},
	}
	for i, c := range cases {
		u, err := createURL(c.in_cmd, c.in_v)
		if err != nil {
			t.Errorf("Error creating url: %v", err)
			continue
		}
		u_str := u.String()
		if u_str != c.want {
			t.Errorf("Case %d. Expected %#v Got %#v", i, c.want, u_str)
		}
	}
}
