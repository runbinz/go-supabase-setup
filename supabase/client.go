package supabase

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Package supabase provides functions to initialize and manage the database connection.
// It uses environment variables to configure the connection to the Supabase Postgres database.
// This package is crucial for setting up the database connection that the application will use to store and retrieve data.

// Init initializes the database connection using the connection string from environment variables.
// It returns an error if the connection fails or the environment variable is not set.
// This function is responsible for establishing a connection to the database, which is essential for any database operations.

func Init() error {
	connStr := os.Getenv("SUPABASE_DB_URL")
	if connStr == "" {
		return fmt.Errorf("SUPABASE_DB_URL is not set")
	}
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to Supabase: %w", err)
	}
	return DB.Ping()
}
