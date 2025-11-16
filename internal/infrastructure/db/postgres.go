package db

import (
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)



func NewPostgres(dsn string) (*sqlx.DB, error) {
    db, err := sqlx.Connect("postgres", dsn)
    if err != nil {
        return nil, err
    }
    // Set sensible defaults (adjust as needed)
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * 60)
    return db, nil
}
