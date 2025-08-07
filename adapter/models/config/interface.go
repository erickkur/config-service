package config

import (
	"context"

	pg "github.com/configservice/adapter/database/sqllite"
)

type ConfigurationModelInterface interface {
	CreateConfiguration(
		dbClient pg.DatabaseAdapterInterface,
		ctx context.Context,
		d Configuration,
	) (*Configuration, error)
}
