package db

import (
    "context"
    "database/sql"
    _ "github.com/lib/pq"
    "os"
)

func Connect() (*sql.DB, error) {
    dsn := os.Getenv("DATABASE_URL")
    return sql.Open("postgres", dsn)
}

func Migrate(db *sql.DB) error {
    sqlBytes, err := os.ReadFile("internal/db/schema.sql")
    if err != nil { return err }
    _, err = db.ExecContext(context.Background(), string(sqlBytes))
    return err
}
