package redisdb

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sifaconer/crud_api/pkg/domain/entity"
)

type StreamRedis struct {
	cli *redis.Client
}

func NewStreamRedis(cli *redis.Client) *StreamRedis {
	return &StreamRedis{
		cli: cli,
	}
}

func (s *StreamRedis) PublishEvent(msg entity.Medidor) error {
	date := ""
	if msg.RetirementDate != nil {
		date = msg.RetirementDate.String()
	}

	err := s.cli.XAdd(context.Background(), &redis.XAddArgs{
		Stream: "Medidor",
		MaxLen: 0,
		ID:     "",
		Values: map[string]interface{}{
			"ID":               msg.ID.String(),
			"Brand":            msg.Brand,
			"Address":          msg.Address,
			"InstallationDate": msg.InstallationDate,
			"RetirementDate":   date,
			"Serial":           msg.Serial,
			"Lines":            msg.Lines,
			"IsActive":         msg.IsActive,
			"CreatedAt":        msg.CreatedAt,
		},
	}).Err()

	return err
}
