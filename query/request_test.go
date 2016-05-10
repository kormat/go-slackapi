package query

import (
	"bitbucket.org/kormaton/slapi/config"
	"net/url"
	"testing"
)

func TestCreateUrl(t *testing.T) {
	// FIXME(kormat): this is a hack, until i figure out how to mock things in Go.
	config.Config.APIEndpoint = "http://api.endpoint"
	cases := []struct {
		in_cmd string
		in_v   url.Values
		want   string
	}{
		{"cmd0", url.Values{}, "http://api.endpoint/cmd0"},
		{"cmd1", url.Values{"p1": []string{"v1"}, "p2": []string{"v2"}},
			"http://api.endpoint/cmd1?p1=v1&p2=v2"},
	}
	for i, c := range cases {
		u := CreateURL(c.in_cmd, c.in_v)
		u_str := u.String()
		if u_str != c.want {
			t.Errorf("Case %d. Expected %#v Got %#v", i, c.want, u_str)
		}
	}
}