package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() error {
	databaseUrl := "postgres://cv_auth_backend:pass-cv1234@db:5432/users"
	var err error
	DB, err = pgxpool.New(context.Background(), databaseUrl)
	return err
}
