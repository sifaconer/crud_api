package db

import (
	"fmt"
	"log"
	"os"

	"github.com/sifaconer/crud_api/pkg/database/postgres"
	"github.com/sifaconer/crud_api/pkg/domain/entity"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type cliGorm struct {
	db *gorm.DB
}

func registerModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Medidor{},
	)
}

func removeModels(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&entity.Medidor{},
	)
}

func NewCliGorm() *cliGorm {
	db := postgres.NewPostgresDB(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSL_MODE"),
	)
	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return &cliGorm{
		db: conn,
	}
}

func (cg *cliGorm) Commands() *cli.Command {
	return &cli.Command{
		Name:  "db",
		Usage: "database migrations",
		Subcommands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "migrate database",
				Action: func(c *cli.Context) error {
					err := registerModels(cg.db)
					if err != nil {
						log.Fatalf("[ERROR]: MIGRATIONS -> %s", err.Error())
					}
					log.Fatalf("[INFO]: MIGRATIONS -> %s", "migrated finish")
					return nil
				},
			},
			{
				Name:  "drop",
				Usage: "drop all tables in database",
				Action: func(c *cli.Context) error {
					err := removeModels(cg.db)
					if err != nil {
						log.Fatalf("[ERROR]: %s", err.Error())
					}
					fmt.Println("drop finish")
					return nil
				},
			},
		},
	}
}
