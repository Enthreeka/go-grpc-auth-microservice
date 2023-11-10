package postgres

import (
	"context"
	"github.com/Enthreeka/auth-microservice/internal/entity"
	"github.com/Enthreeka/auth-microservice/internal/repo"
	"github.com/Enthreeka/auth-microservice/pkg/relationDB"
)

type tokenPostgresRepo struct {
	pool relationDB.Pool
}

func NewTokenPostgresRepo(pool relationDB.Pool) repo.TokenRepository {
	return &tokenPostgresRepo{
		pool: pool,
	}
}

func (t *tokenPostgresRepo) GetTokenByID(ctx context.Context, id string) (*entity.Token, error) {
	query := `SELECT id,user_id,refresh_token,expires_at FROM token WHERE id = $1`
	var token entity.Token

	err := t.pool.QueryRow(ctx, query, id).Scan(&token.ID, &token.UserID, &token.RefreshToken, &token.ExpiresAt)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (t *tokenPostgresRepo) DeleteTokenByID(ctx context.Context, id string) error {
	query := `DELETE FROM token WHERE id = $1`

	_, err := t.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *tokenPostgresRepo) CreateToken(ctx context.Context, token *entity.Token) (*entity.Token, error) {
	query := `INSERT INTO token (user_id, refresh_token, expires_at) VALUES ($1,$2,$3) RETURNING id,user_id, refresh_token, expires_at`

	createdToken := new(entity.Token)
	err := t.pool.QueryRow(ctx, query, token.UserID, token.RefreshToken, token.ExpiresAt).Scan(
		&createdToken.ID,
		&createdToken.UserID,
		&createdToken.RefreshToken,
		&createdToken.ExpiresAt)
	if err != nil {
		return nil, err
	}

	return createdToken, nil
}
