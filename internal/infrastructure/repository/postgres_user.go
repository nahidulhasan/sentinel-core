package repository

import (
    "github.com/jmoiron/sqlx"
    "github.com/nahidulhasan/sentinel-core/internal/domain/user"
)

type PostgresUserRepo struct {
    db *sqlx.DB
}

func NewPostgresUserRepo(db *sqlx.DB) user.Repository {
    return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) Create(u *user.User) error {
    query := `INSERT INTO users (name, email) VALUES ($1, $2)`
    _, err := r.db.Exec(query, u.Name, u.Email)
    return err
}

func (r *PostgresUserRepo) List() ([]*user.User, error) {
    users := []*user.User{}
    query := `SELECT id, name, email FROM users`

    err := r.db.Select(&users, query)
    if err != nil {
        return nil, err
    }

    return users, nil
}
