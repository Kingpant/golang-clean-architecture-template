package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Kingpant/golang-template/cmd/bun/migrations"
	"github.com/Kingpant/golang-template/internal/infrastructure/config"
	"github.com/Kingpant/golang-template/internal/infrastructure/db"
	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "bun",
		Commands: []*cli.Command{
			newDBCommand(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newDBCommand() *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database migrations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				FilePath: "etc/db.yml",
				Usage:    "Load configuration from `FILE`",
			},
			&cli.StringFlag{
				Name:     "seed",
				Aliases:  []string{"s"},
				FilePath: "etc/seed.yml",
				Usage:    "Load seeders from `FILE`",
			},
		},
		Before: func(c *cli.Context) error {
			cfg, loadErr := config.LoadMigrationConfig()
			if loadErr != nil {
				panic(loadErr)
			}
			db := db.NewDB(
				cfg.AppEnv,
				cfg.PostgresqlUsername,
				cfg.PostgresqlPassword,
				cfg.PostgresqlHost,
				cfg.PostgresqlDatabase,
				cfg.PostgresqlSchema,
				cfg.PostgresqlSSL,
			)

			c.Context = context.WithValue(
				c.Context,
				"db-migrator",
				migrate.NewMigrator(db, migrations.Migrations),
			)
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:  "init",
				Usage: "create migration tables",
				Action: func(c *cli.Context) error {
					migrator := c.Context.Value("db-migrator").(*migrate.Migrator)
					return migrator.Init(c.Context)
				},
			},
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					migrator := c.Context.Value("db-migrator").(*migrate.Migrator)

					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer migrator.Unlock(c.Context)

					group, err := migrator.Migrate(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to run (database is up to date)\n")
						return nil
					}
					fmt.Printf("migrated to %s\n", group)
					return nil
				},
			},
			{
				Name:  "rollback",
				Usage: "rollback the last migration group",
				Action: func(c *cli.Context) error {
					migrator := c.Context.Value("db-migrator").(*migrate.Migrator)

					if err := migrator.Lock(c.Context); err != nil {
						return err
					}
					defer migrator.Unlock(c.Context)

					group, err := migrator.Rollback(c.Context)
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no groups to roll back\n")
						return nil
					}
					fmt.Printf("rolled back %s\n", group)
					return nil
				},
			},
			{
				Name:  "lock",
				Usage: "lock migrations",
				Action: func(c *cli.Context) error {
					migrator := c.Context.Value("db-migrator").(*migrate.Migrator)
					return migrator.Lock(c.Context)
				},
			},
			{
				Name:  "unlock",
				Usage: "unlock migrations",
				Action: func(c *cli.Context) error {
					migrator := c.Context.Value("db-migrator").(*migrate.Migrator)
					return migrator.Unlock(c.Context)
				},
			},
			{
				Name:  "create_go",
				Usage: "create Go migration",
				Action: func(c *cli.Context) error {
					migrator := c.Context.Value("db-migrator").(*migrate.Migrator)

					name := strings.Join(c.Args().Slice(), "_")
					mf, err := migrator.CreateGoMigration(c.Context, name)
					if err != nil {
						return err
					}
					fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					return nil
				},
			},
			{
				Name:  "create_sql",
				Usage: "create up and down SQL migrations",
				Action: func(c *cli.Context) error {
					migrator := c.Context.Value("db-migrator").(*migrate.Migrator)

					name := strings.Join(c.Args().Slice(), "_")
					files, err := migrator.CreateSQLMigrations(c.Context, name)
					if err != nil {
						return err
					}

					for _, mf := range files {
						fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
					}

					return nil
				},
			},
			{
				Name:  "status",
				Usage: "print migrations status",
				Action: func(c *cli.Context) error {
					migrator := c.Context.Value("db-migrator").(*migrate.Migrator)

					ms, err := migrator.MigrationsWithStatus(c.Context)
					if err != nil {
						return err
					}
					fmt.Printf("migrations: %s\n", ms)
					fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
					fmt.Printf("last migration group: %s\n", ms.LastGroup())
					return nil
				},
			},
			{
				Name:  "mark_applied",
				Usage: "mark migrations as applied without actually running them",
				Action: func(c *cli.Context) error {
					migrator := c.Context.Value("db-migrator").(*migrate.Migrator)

					group, err := migrator.Migrate(c.Context, migrate.WithNopMigration())
					if err != nil {
						return err
					}
					if group.IsZero() {
						fmt.Printf("there are no new migrations to mark as applied\n")
						return nil
					}
					fmt.Printf("marked as applied %s\n", group)
					return nil
				},
			},
		},
	}
}
