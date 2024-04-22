package ginToolkit

import (
	"github.com/gin-gonic/gin"
)

type BusinessLogicFunc func(engine *gin.Engine) error

func CreateEngine(c *Config, businessLogic BusinessLogicFunc, options ...Option) {
	// opts := loadOptions(options...)

}

func loadOptions(options ...Option) *option {
	opts := &option{
		ServerInfo:            "",
		RecoveryMiddleware:    gin.Recovery(),
		UseDefaultNoRouteHtml: true,
		UseDefaultNoMethod:    true,
		UseDefaultFavicon:     true,
	}
	for _, opt := range options {
		opt.apply(opts)
	}
	return opts
}
