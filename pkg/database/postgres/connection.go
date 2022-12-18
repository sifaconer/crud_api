package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	user     string
	password string
	host     string
	port     string
	db       string
	ssl      string
}

func NewPostgresDB(user, password, host, port, db, ssl string) *postgresDB {
	return &postgresDB{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		db:       db,
		ssl:      ssl,
	}
}

func (p *postgresDB) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.user,
		p.password,
		p.host,
		p.port,
		p.db,
		p.ssl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
