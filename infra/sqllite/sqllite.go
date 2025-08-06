package sqllite

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/url"

	"github.com/configservice/internal/env"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Database struct {
	Conn *bun.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) Connect() {
	var tlsConfig *tls.Config

	parsedURL, err := url.Parse(env.DBUrl())
	if err != nil {
		log.Fatalf("Error parsing database URL: %v", err)
	}

	connector := pgdriver.NewConnector(
		pgdriver.WithApplicationName(env.AppName()),
		pgdriver.WithDSN(parsedURL.String()),
		pgdriver.WithTLSConfig(tlsConfig),
	)

	db.Conn = bun.NewDB(sql.OpenDB(connector), pgdialect.New())
}

func (db *Database) Close() {
	db.Conn.Close()
}
