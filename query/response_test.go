package query

import (
	"fmt"
	"testing"
)

func TestParseResponse(t *testing.T) {
	cases := []struct {
		in                 []byte
		ok                 bool
		warn, err, channel *string
	}{
		{in: []byte(`{"ok": true, "channel": {"name": "This is good"}}`),
			ok: true, channel: stringPtr(`{"name": "This is good"}`)},
		{in: []byte(`{"ok": false, "error": "channel_not_found"}`),
			ok: false, err: stringPtr("channel_not_found")},
		{in: []byte(`{"ok": true, "warning": "something_problematic", "channel": {"name": "Your requested information"}}`),
			ok: true, warn: stringPtr("something_problematic"),
			channel: stringPtr(`{"name": "Your requested information"}`)},
	}
	for i, c := range cases {
		pass := true
		got, _ := Parse(c.in)
		if c.ok != got.Ok {
			t.Errorf("Case %d. Ok: Expected %v Got %v", i, c.ok, got.Ok)
			pass = false
		}
		if !cmpStrPtr(c.warn, got.Warning) {
			t.Errorf("Case %d. Warning: Expected %v Got %v", i, fmtStrPtr(c.warn), fmtStrPtr(got.Warning))
			pass = false
		}
		if !cmpStrPtr(c.err, got.Error) {
			t.Errorf("Case %d. Error: Expected %v Got %v", i, fmtStrPtr(c.err), fmtStrPtr(got.Error))
			pass = false
		}
		var chnptr *string
		if c.channel != nil {
			chnptr = stringPtr(string(*got.Channel))
		}
		if !cmpStrPtr(c.channel, chnptr) {
			t.Errorf("Case %d. Raw channel: Expected %v Got %v", i, fmtStrPtr(c.channel), fmtStrPtr(chnptr))
			pass = false
		}
		if !pass {
			t.Log(got.String())
		}
	}
}

func fmtStrPtr(s *string) string {
	if s == nil {
		return "nil"
	}
	return fmt.Sprintf("`%s`", *s)
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
