package main

import (
	"log"
	"os"

	"github.com/sifaconer/crud_api/cmd/server/cligrpc"
	"github.com/sifaconer/crud_api/cmd/server/db"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "crud_api",
		Usage: "Cli to interact app",
		Commands: []*cli.Command{
			db.NewCliGorm().Commands(),
			cligrpc.NewCliGRPC().Commands(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func init() {
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PORT := os.Getenv("POSTGRES_PORT")
	POSTGRES_HOST := os.Getenv("POSTGRES_HOST")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	POSTGRES_SSL_MODE := os.Getenv("POSTGRES_SSL_MODE")
	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PORT := os.Getenv("REDIS_PORT")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")
	SERVER_PORT := os.Getenv("SERVER_PORT")
	valid := func(name, content string) {
		if len(content) == 0 {
			log.Fatalf("[ERROR]: environment %s is empty\n", name)
		}
	}

	valid("POSTGRES_USER", POSTGRES_USER)
	valid("POSTGRES_PORT", POSTGRES_PORT)
	valid("POSTGRES_HOST", POSTGRES_HOST)
	valid("POSTGRES_PASSWORD", POSTGRES_PASSWORD)
	valid("POSTGRES_DB", POSTGRES_DB)
	valid("POSTGRES_SSL_MODE", POSTGRES_SSL_MODE)
	valid("REDIS_HOST", REDIS_HOST)
	valid("REDIS_PORT", REDIS_PORT)
	valid("REDIS_PASSWORD", REDIS_PASSWORD)
	valid("SERVER_PORT", SERVER_PORT)
}
