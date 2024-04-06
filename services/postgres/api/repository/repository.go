package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func NewClient() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_ADDR"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB")))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
