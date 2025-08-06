package infra

import (
	"github.com/configservice/infra/sqllite"
)

type Infra struct {
	Database *sqllite.Database
}

func Init() *Infra {
	database := sqllite.NewDatabase()

	return &Infra{
		Database: database,
	}
}
