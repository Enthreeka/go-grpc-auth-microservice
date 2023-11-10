package redis

import (
	"context"
	"github.com/Enthreeka/auth-microservice/internal/entity"
	"github.com/Enthreeka/auth-microservice/internal/repo"
	pkg "github.com/Enthreeka/auth-microservice/pkg/redis"
	"reflect"
	"sync"
)

type tokenRedisRepo struct {
	*pkg.Redis

	sync.Mutex
}

func NewTokenRedisRepo(redis *pkg.Redis) repo.TokenRepository {
	return &tokenRedisRepo{
		Redis: redis,
	}
}

func (t *tokenRedisRepo) GetTokenByID(ctx context.Context, id string) (*entity.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tokenRedisRepo) DeleteTokenByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (t *tokenRedisRepo) CreateToken(ctx context.Context, token *entity.Token) (*entity.Token, error) {
	//bytesToken, err := json.Marshal(token)
	//if err != nil {
	//	return nil, err
	//}

	//err := t.Rds.HSet(ctx, token.UserID.String(), map[string]interface{}{
	//	"expires_at":    token.ExpiresAt,
	//	"refresh_token": token.RefreshToken,
	//}).Err()
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = t.Rds.Set(ctx, token.UserID.String(), bytesToken, 360*time.Hour).Err()
	//if err != nil {
	//	return nil, err
	//}

	valueStructToken := reflect.ValueOf(token)
	for i := 0; i < valueStructToken.NumField(); i++ {
		field := valueStructToken.Field(i)
		//fieldName := valueStructToken.Type().Field(i).Name

		err := t.Rds.Append(ctx, token.UserID.String(), field.String()).Err()
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
