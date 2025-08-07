package configuration

import (
	cService "github.com/configservice/application/services/configuration"
	rs "github.com/configservice/application/services/router"
	"github.com/configservice/internal/constant"
	"github.com/configservice/internal/logger"
)

type RouteDependency struct {
	Context              rs.Context
	Logger               logger.Interface
	ConfigurationService cService.ConfigurationServiceInterface
}

type Route struct {
	Context              rs.Context
	Logger               logger.Interface
	ConfigurationService cService.ConfigurationServiceInterface
}

func NewDomain(d RouteDependency) {
	route := Route(d)

	route.initEndpoints()
}

func (r Route) initEndpoints() {
	h := NewHandler(HandlerDependency{
		Logger:               r.Logger,
		Context:              r.Context,
		ConfigurationService: r.ConfigurationService,
	})

	r.Context.RegisterEndpoint(r.CreateConfigurationEndpoint(h))
}

func (r Route) CreateConfigurationEndpoint(h Handler) rs.EndpointInfo {
	return rs.EndpointInfo{
		HTTPMethod: "POST",
		URLPattern: "/configurations",
		Handler:    h.CreateConfigurationHandler(),
		Verifications: []constant.VerificationType{
			constant.VerificationTypeConstants.AppToken,
		},
	}
}
