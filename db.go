package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func conn() (*pgx.Conn, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), ("postgres://postgres:password@localhost:5432/postgres"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connection was successful")
	return conn, err
}

/*
docker run --name some-postgres \
  -e POSTGRES_PASSWORD=password \
  -p 5432:5432 \
  -d postgres


*/
