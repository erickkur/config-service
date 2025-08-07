package configuration

import (
	"context"

	"github.com/configservice/application/dto"
	cService "github.com/configservice/application/services/configuration"
	errs "github.com/configservice/internal/error"
)

type EntityDependency struct {
	ConfigurationService cService.ConfigurationServiceInterface
}

type Entity struct {
	configurationService cService.ConfigurationServiceInterface
}

func NewEntity(dependency EntityDependency) Entity {
	return Entity{
		configurationService: dependency.ConfigurationService,
	}
}

func (e Entity) CreateConfiguration(ctx context.Context, request dto.CreateConfigurationRequest) (*dto.CreateConfigurationResponse, *errs.JSONWrapError) {
	return nil, nil
}
