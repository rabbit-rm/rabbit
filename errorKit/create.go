package errorKit

import (
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/rabbit-rm/rabbit/stackKit"
)

const skip = 1

// New 创建一个新的自定义错误，包含堆栈信息
func New(format string, args ...interface{}) error {
	return gerror.NewSkipf(skip, format, args...)
}

func NewWithCaller(format string, args ...interface{}) error {
	return New(addCaller(format), args...)
}

func addCaller(format string) string {
	return fmt.Sprintf("%s -> %s", stackKit.Caller(skip), format)
}

// Wrap 包裹其他错误，用于构造多级错误，包含堆栈信息
func Wrap(err error, format string, args ...interface{}) error {
	return gerror.WrapSkipf(skip, err, format, args...)
}

func WrapWithCaller(err error, format string, args ...interface{}) error {
	return Wrap(err, addCaller(format), args...)
}
