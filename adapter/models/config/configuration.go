package config

import (
	"context"

	pg "github.com/configservice/adapter/database/sqllite"
	"github.com/configservice/adapter/models/base"
	"github.com/configservice/adapter/models/recordtimestamp"
	"github.com/uptrace/bun"
)

type Configuration struct {
	base.Base
	recordtimestamp.RecordTimestamp
	Name    string `bun:",notnull"`
	Version int32  `bun:",notnull"`
	Data    string `bun:",notnull"`

	bun.BaseModel `bun:"configurations,alias:c"`
}

type ConfigurationModel struct {
}

func NewModel() *ConfigurationModel {
	return &ConfigurationModel{}
}

func (di *ConfigurationModel) CreateConfiguration(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	d Configuration,
) (*Configuration, error) {
	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	query := db.GetConnectionDB().
		NewInsert().
		Model(&d)
	_, err = query.Exec(ctx)

	return &d, err
}
