package query

import (
	"net/url"
	"testing"
)

func TestCreateUrl(t *testing.T) {
	cases := []struct {
		in_cmd string
		in_v   url.Values
		want   string
	}{
		{"cmd0", url.Values{}, "https://slack.com/api/cmd0"},
		{"cmd1", url.Values{"p1": []string{"v1"}, "p2": []string{"v2"}},
			"https://slack.com/api/cmd1?p1=v1&p2=v2"},
	}
	for i, c := range cases {
		u := CreateURL(c.in_cmd, c.in_v)
		u_str := u.String()
		if u_str != c.want {
			t.Errorf("Case %d. Expected %#v Got %#v", i, c.want, u_str)
		}
	}
}
