package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Conn *pgxpool.Pool

func NewConnection() {
	getenv := os.Getenv("DATABASE_URL")
	if len(getenv) == 0 {
		getenv = "postgresql://postgres:secret@localhost:5433/postgres?sslmode=disable"
	}
	dbPool, err := pgxpool.Connect(context.Background(), getenv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Conn = dbPool
}

func CloseConnection() {
	defer Conn.Close()
}
