package test

import (
	"fmt"
	"reflect"
	"testing"
)

type TestMeta struct {
	T  *testing.T
	Ok bool
}

func (tm *TestMeta) Eq(name string, want, got interface{}) {
	var type_fmt string
	var equal = false
	// Determine the type of comparison required
	switch want.(type) {
	default:
		equal = want == got
	case []string:
		equal = reflect.DeepEqual(want, got)
	}
	if !equal {
		tm.Ok = false
		/// Determine the formatting for the value type
		switch want.(type) {
		default:
			type_fmt = "%v"
		case string:
			type_fmt = "`%s`"
		}
		tm.T.Errorf("%s: Expected %s Got %s", name, fmt.Sprintf(type_fmt, want), fmt.Sprintf(type_fmt, got))
	}
}
