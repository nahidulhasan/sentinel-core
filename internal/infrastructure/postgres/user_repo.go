package userrepo

import (
    "github.com/jmoiron/sqlx"
    "github.com/nahidulhasan/sentinel-core/internal/domain/user"
)

type PostgresRepo struct {
    db *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) *PostgresRepo {
    return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Create(u *user.User) error {
    query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
    return r.db.QueryRow(query, u.Name, u.Email).Scan(&u.ID)
}

func (r *PostgresRepo) List() ([]*user.User, error) {
    users := []*user.User{}
    err := r.db.Select(&users, `SELECT id, name, email FROM users ORDER BY id DESC`)
    return users, err
}
