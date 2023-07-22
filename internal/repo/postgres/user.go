package postgres

import (
	"context"
	"github.com/NASandGAP/auth-microservice/internal/entity"
	"github.com/NASandGAP/auth-microservice/internal/repo"
	"github.com/NASandGAP/auth-microservice/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userPostgresRepo struct {
	pool *pgxpool.Pool
	log  *logger.Logger
}

func NewUserPostgresRepo(pool *pgxpool.Pool, log *logger.Logger) repo.Repository {
	return &userPostgresRepo{
		pool: pool,
		log:  log,
	}
}

func (u *userPostgresRepo) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	query := "SELECT id, email , password FROM user WHERE id = $1"
	var user entity.User

	err := u.pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userPostgresRepo) DeleteUserByID(ctx context.Context, id string) error {
	query := "DELETE FROM user WHERE id = &1"

	_, err := u.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userPostgresRepo) CreateUser(ctx context.Context, user entity.User) error {
	query := "INSERT INTO user (email,password) VALUES ($1,$2)"

	_, err := u.pool.Exec(ctx, query, &user.Email, &user.Password)
	if err != nil {
		return err
	}
	return nil
}
