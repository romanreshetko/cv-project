package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateUser(username, passwordHash string) error {
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2)`
	_, err := DB.Exec(context.Background(), query, username, passwordHash)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return errors.New("username already exists")
		}
		return err
	}
	return nil
}
