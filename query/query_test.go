package query

import (
	"fmt"
	"testing"
)

func TestParseResponse(t *testing.T) {
	cases := []struct {
		in           []byte
		ok           bool
		warning, err *string
		channel      bool
	}{
		{in: []byte(`{"ok": true, "channel": {"name": "This is good"}}`),
			ok: true, channel: true},
		{in: []byte(`{"ok": false, "error": "channel_not_found"}`),
			ok: false, err: stringPtr("channel_not_found")},
		{in: []byte(`{"ok": true, "warning": "something_problematic", "channel": {"name": "Your requested information"}}`),
			ok: true, warning: stringPtr("something_problematic"), channel: true},
	}
	pass := true
	for i, c := range cases {
		got, ok := parse(c.in)
		if !ok {
			t.Errorf("Case %d. Error parsing json", i)
			pass = false
		}
		if got.Ok != c.ok {
			t.Errorf("Case %d. Ok: Expected %#v Got %#v", i, c.ok, got.Ok)
			pass = false
		}
		if !cmpStrPtr(c.warning, got.Warning) {
			t.Errorf("Case %d. Warning: Expected %#v Got %#v", i, fmtStr(c.warning), fmtStr(got.Warning))
			pass = false
		}
		if !cmpStrPtr(c.err, got.Error) {
			t.Errorf("Case %d. Error: Expected %v Got %v", i, fmtStr(c.err), fmtStr(got.Error))
			pass = false
		}
		has_chan := got.Channel != nil
		if has_chan != c.channel {
			t.Errorf("Case %d. Channel presence: Expected %#v Got %#v", i, c.channel, has_chan)
			pass = false
		}
		if !pass {
			t.Log(got.String())
		}
	}
}

func fmtStr(s *string) string {
	if s == nil {
		return "nil"
	}
	return fmt.Sprintf("%#v", *s)
}

func cmpStrPtr(want, got *string) bool {
	if want == nil {
		return got == nil
	}
	if got == nil {
		return false
	}
	return *want == *got
}

func stringPtr(s string) *string { return &s }
