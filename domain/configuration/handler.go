package configuration

import (
	"net/http"

	"github.com/configservice/application/dto"
	cService "github.com/configservice/application/services/configuration"
	rs "github.com/configservice/application/services/router"
	errs "github.com/configservice/internal/error"
	"github.com/configservice/internal/handler"
	"github.com/configservice/internal/json"
	"github.com/configservice/internal/logger"
)

type HandlerDependency struct {
	Logger               logger.Interface
	Context              rs.Context
	ConfigurationService cService.ConfigurationServiceInterface
}

type Handler struct {
	logger               logger.Interface
	context              rs.Context
	configurationService cService.ConfigurationServiceInterface
	resp                 handler.ResponseInterface
	entity               Entity
}

func NewHandler(d HandlerDependency) Handler {
	entity := NewEntity(EntityDependency{
		ConfigurationService: d.ConfigurationService,
	})

	return Handler{
		logger:               d.Logger,
		context:              d.Context,
		configurationService: d.ConfigurationService,
		resp:                 handler.NewResponse(handler.Dep{}),
		entity:               entity,
	}
}

func (h Handler) CreateConfigurationHandler() handler.EndpointHandler {
	return func(w http.ResponseWriter, r *http.Request) handler.ResponseInterface {
		ctx := r.Context()

		var request dto.CreateConfigurationRequest
		err := json.DecodeBody(&request, r.Body)
		if err != nil {
			decodingErr := err.WrapError(errs.ConfigurationPrefix)
			return h.resp.ImportJSONWrapError(&decodingErr)
		}

		response, jsonWrapError := h.entity.CreateConfiguration(ctx, request)
		if jsonWrapError != nil {
			return h.resp.ImportJSONWrapError(jsonWrapError)
		}

		return h.resp.SetOkWithStatus(http.StatusCreated, response)
	}
}
