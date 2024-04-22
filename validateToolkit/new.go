package validateToolkit

import (
	"github.com/go-playground/validator/v10"
)

// New 返回 validator.Validate 实例
// 不传参，默认使用 validate tag，传参使用第一个作为 tag name
func New(tagNames ...string) *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())
	if len(tagNames) > 0 {
		v.SetTagName(tagNames[0])
	}
	return v
}
