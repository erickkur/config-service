package router

import (
	"path"

	"github.com/configservice/adapter/middleware"
	routerAdapter "github.com/configservice/adapter/router"
	"github.com/configservice/internal/constant"
	"github.com/configservice/internal/handler"
	"github.com/thoas/go-funk"
)

type Context struct {
	router     routerAdapter.RouterInterface
	middleware *middleware.Middleware
	prefix     string
}

type EndpointInfo struct {
	HTTPMethod    string
	URLPattern    string
	Handler       handler.EndpointHandler
	Verifications []constant.VerificationType
}

func NewService(
	router routerAdapter.RouterInterface,
	middleware middleware.Middleware,
	prefix string,
) Context {
	return Context{
		router:     router,
		middleware: &middleware,
		prefix:     prefix,
	}
}

// RegisterEndpoint ...
func (r *Context) RegisterEndpoint(info EndpointInfo) {
	r.RegisterEndpointWithPrefix(info, r.prefix)
}

// RegisterEndpointWithPrefix ...
func (r *Context) RegisterEndpointWithPrefix(info EndpointInfo, prefix string) {
	m := r.middleware
	urlPattern := getFullURLPattern(info, prefix)
	verificationFns := getVerificationMethod(m, info.Verifications)

	r.router.Handle(info.HTTPMethod, urlPattern, m.Verify(info.Handler, verificationFns...))
}

func getVerificationMethod(m *middleware.Middleware, verifications []constant.VerificationType) []middleware.MiddlewareFunc {
	return funk.Map(verifications, func(_ constant.VerificationType) middleware.MiddlewareFunc {
		return m.AppToken
	}).([]middleware.MiddlewareFunc)
}

func getFullURLPattern(info EndpointInfo, prefix string) string {
	if prefix == "" {
		return info.URLPattern
	}

	return path.Join(prefix, info.URLPattern)
}
