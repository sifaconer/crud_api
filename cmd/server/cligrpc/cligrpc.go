package cligrpc

import (
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/sifaconer/crud_api/pkg/database/postgres"
	"github.com/sifaconer/crud_api/pkg/database/postgres/query"
	"github.com/sifaconer/crud_api/pkg/database/redisdb"
	"github.com/sifaconer/crud_api/pkg/domain/usecase"
	"github.com/sifaconer/crud_api/pkg/grpc"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
)

type cliGRPC struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewCliGRPC() *cliGRPC {
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

	rdb := redisdb.NewRedisDB(
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"))
	cli, err := rdb.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return &cliGRPC{
		db:  conn,
		rdb: cli,
	}
}

func (cg *cliGRPC) Commands() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Server http",
		Subcommands: []*cli.Command{
			{
				Name:  "grpc",
				Usage: "run grpc server",
				Action: func(c *cli.Context) error {
					port := os.Getenv("SERVER_PORT")
					log.Printf("[INFO]: Run Server in http://127.0.0.1:%s", port)

					rcli := redisdb.NewStreamRedis(cg.rdb)

					medidorQuery := query.NewMedidorQuery(cg.db)
					medidorUsecase := usecase.NewMedidorUseCaseImpl(
						medidorQuery, rcli)

					server := grpc.NewServerGRPC(
						port,
						medidorUsecase,
					)

					if err := server.RunServer(); err != nil {
						log.Fatalf("[ERROR]: %s", err.Error())
					}
					return nil
				},
			},
		},
	}
}
