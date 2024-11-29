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

func (rep *RepositoryUser) GetUserByUsernameOrEmail(ctx context.Context, email string, username string) (User, error) {
	query := `SELECT id, email, username, password, created_at FROM users WHERE email = $1 OR username = $2`
	row := rep.DB.QueryRow(ctx, query, email, username)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt)
	return user, err
}

func (rep *RepositoryUser) GetUserByID(ctx context.Context, userID int) (*User, error) {
	query := `SELECT id, email, username, password, created_at FROM users WHERE id = $1`
	row := rep.DB.QueryRow(ctx, query, userID)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt)
	return &user, err
}
