package errorKit

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/errors/gerror"
)

func TestNew(t *testing.T) {
	var a int
	err := New("invalid data type(%T)", a)
	if err != nil {
		fmt.Printf("error:%+v\n", err)
		json, err := err.(*gerror.Error).MarshalJSON()
		if err != nil {
			return
		}
		fmt.Printf("error:%s\n", json)
	}
}

func TestNewWithCaller(t *testing.T) {
	var a int
	err := NewWithCaller("invalid data type(%T)", a)
	if err != nil {
		fmt.Printf("error:%+v\n", err)
		json, err := err.(*gerror.Error).MarshalJSON()
		if err != nil {
			return
		}
		fmt.Printf("error:%s\n", json)
	}
}

func TestWrap(t *testing.T) {
	var err = errors.New("invalid data")
	err = Wrap(err, "value:%v", 12)
	if err != nil {
		fmt.Printf("error:%+v\n", err)
	}
}

func TestWrapWithCaller(t *testing.T) {
	var err = errors.New("invalid data")
	err = WrapWithCaller(err, "value:%v", 12)
	if err != nil {
		fmt.Printf("error:%+v\n", err)
	}
}
