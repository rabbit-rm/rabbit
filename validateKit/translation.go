package validateKit

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"
)

const (
	zhTrans = "zh"
	enTrans = "en"
)

var (
	uni *ut.UniversalTranslator
)

func init() {
	var (
		zht = zh.New()
		ent = en.New()
	)
	uni = ut.New(zht, zht, ent)
	// register
	panicIf(zhTranslation.RegisterDefaultTranslations(validate, ZhTrans()))
	panicIf(zhTranslation.RegisterDefaultTranslations(validate, EnTrans()))
}

func trans(local string) ut.Translator {
	t, found := uni.GetTranslator(local)
	if found {
		return t
	}
	return nil
}

func ZhTrans() ut.Translator {
	return trans(zhTrans)
}
func EnTrans() ut.Translator {
	return trans(enTrans)
}
