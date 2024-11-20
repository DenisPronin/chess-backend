package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryUser struct {
	DB *pgxpool.Pool
}

func NewRepositoryUser(db *pgxpool.Pool) *RepositoryUser {
	return &RepositoryUser{DB: db}
}

func (rep *RepositoryUser) SaveUser(ctx context.Context, user User) error {
	query := `INSERT INTO users (email, username, password, created_at) VALUES ($1, $2, $3, $4)`
	_, err := rep.DB.Exec(ctx, query, user.Email, user.Username, user.Password, user.CreatedAt)
	return err
}
