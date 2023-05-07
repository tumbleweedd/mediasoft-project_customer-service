package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	PgPort         string
	PgHost         string
	PgDBName       string
	PgDBSchemaName string
	PgUser         string
	PgPwd          string
	PgSSLMode      string
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s search_path=%s",
			cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgDBName, cfg.PgPwd, cfg.PgSSLMode, cfg.PgDBSchemaName),
	)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
