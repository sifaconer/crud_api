package redisdb

import "github.com/go-redis/redis/v8"

type redisDB struct {
	password string
	port     string
	host     string
}

func NewRedisDB(password, host, port string) *redisDB {
	return &redisDB{
		password: password,
		port:     port,
		host:     host,
	}
}

func (p *redisDB) Connect() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: p.host + ":" + p.port,
		//Password: p.password,
		DB: 0,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
