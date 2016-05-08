package channels

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParseList(t *testing.T) {
	cases := []struct {
		in   []byte
		want []json.RawMessage
	}{
		{[]byte(`[{"name": "chan1"}]`), []json.RawMessage{json.RawMessage(`{"name": "chan1"}`)}},
		{[]byte(`[{"name": "chan1"},{"name": "chan2"}]`),
			[]json.RawMessage{json.RawMessage(`{"name": "chan1"}`), json.RawMessage(`{"name": "chan2"}`)}},
	}
	for i, c := range cases {
		cl, ok := parseList(c.in)
		if !ok {
			t.Errorf("Case %d. Error parsing json", i)
		}
		if !reflect.DeepEqual(c.want, cl.Channels) {
			t.Errorf("Case %d. Expected %s Got %s", i, c.want, cl.Channels)
		}
	}
}
