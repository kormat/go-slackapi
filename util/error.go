package util

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
)

func Error(format string, a ...interface{}) error {
	errstr := fmt.Sprintf(format, a...)
	glog.Error(errstr)
	return errors.New(errstr)
}
