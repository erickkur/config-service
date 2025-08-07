package configuration

import (
	"context"

	pg "github.com/configservice/adapter/database/sqllite"
	configurationAdapter "github.com/configservice/adapter/models/config"
	"github.com/configservice/application/dto"
)

type Dependency struct {
	ConfigurationModel configurationAdapter.ConfigurationModelInterface
	DBClient           pg.DatabaseAdapterInterface
}

type ConfigurationService struct {
	configurationModel configurationAdapter.ConfigurationModelInterface
	dbClient           pg.DatabaseAdapterInterface
}

func NewConfigurationService(dependency Dependency) *ConfigurationService {
	return &ConfigurationService{
		configurationModel: dependency.ConfigurationModel,
		dbClient:           dependency.DBClient,
	}
}

func (ds *ConfigurationService) CreateConfiguration(ctx context.Context, request dto.CreateConfigurationRequest) (*configurationAdapter.Configuration, error) {
	return nil, nil
}
