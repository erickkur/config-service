package configuration

import (
	"context"

	configurationAdapter "github.com/configservice/adapter/models/configuration"
	"github.com/configservice/application/dto"
)

type ConfigurationServiceInterface interface {
	CreateConfiguration(ctx context.Context, request dto.CreateConfigurationRequest) (*configurationAdapter.Configuration, error)
}
