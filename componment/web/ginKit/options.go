package ginKit

import (
	"github.com/gin-gonic/gin"
)

type Option interface {
	apply(*option)
}

type optionFunc func(option *option)

func (f optionFunc) apply(option *option) {
	f(option)
}

type option struct {
	ServerInfo            string
	RecoveryMiddleware    gin.HandlerFunc
	UseDefaultNoRouteHtml bool
	UseDefaultNoMethod    bool
	UseDefaultFavicon     bool
}

func WithServiceInfo(info string) Option {
	return optionFunc(func(option *option) {
		option.ServerInfo = info
	})
}

func WithRecoveryMiddleware(recovery gin.HandlerFunc) Option {
	return optionFunc(func(option *option) {
		option.RecoveryMiddleware = recovery
	})
}

func UseDefaultNoRouteHtml(use bool) Option {
	return optionFunc(func(option *option) {
		option.UseDefaultNoRouteHtml = use
	})
}

func UseDefaultNoMethod(use bool) Option {
	return optionFunc(func(option *option) {
		option.UseDefaultNoMethod = use
	})
}

func UseDefaultFavicon(use bool) Option {
	return optionFunc(func(option *option) {
		option.UseDefaultFavicon = use
	})
}
