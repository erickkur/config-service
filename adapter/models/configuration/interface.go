package configuration

import (
	"context"

	sl "github.com/configservice/adapter/database/sqllite"
)

type ConfigurationModelInterface interface {
	CreateConfiguration(
		dbClient sl.DatabaseAdapterInterface,
		ctx context.Context,
		d Configuration,
	) (*Configuration, error)
}
