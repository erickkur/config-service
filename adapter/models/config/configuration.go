package configuration

import (
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
