package db

import (
	"database/sql"
	"net/url"
	"time"

	"github.com/Kingpant/golang-clean-architecture-template/internal/infrastructure/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewDB(
	appEnv config.AppEnvType,
	userName, password, host, database, schema string,
	ssl bool,
) *bun.DB {
	connectionUrl := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(userName, password),
		Host:   host,
		Path:   database,
	}
	connectionString := connectionUrl.String()
	pgOpts := []pgdriver.Option{
		pgdriver.WithDSN(connectionString),
		pgdriver.WithConnParams(map[string]any{
			"search_path": schema,
		}),
	}
	if !ssl {
		pgOpts = append(pgOpts, pgdriver.WithInsecure(true))
	}
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgOpts...))

	sqldb.SetMaxOpenConns(50)
	sqldb.SetMaxIdleConns(25)
	sqldb.SetConnMaxIdleTime(10 * time.Minute)
	sqldb.SetConnMaxLifetime(1 * time.Hour)

	db := bun.NewDB(sqldb, pgdialect.New())

	if appEnv == config.AppEnvLocal {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))
	}

	return db
}
