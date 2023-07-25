package postgres

import (
	"context"
	"github.com/NASandGAP/auth-microservice/internal/entity"
	"github.com/google/uuid"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserPostgresRepo_CreateUser(t *testing.T) {
	t.Parallel()

	t.Run("Testing create user in Postgres", func(t *testing.T) {

		mockDB, err := pgxmock.NewPool()
		require.NoError(t, err)
		defer mockDB.Close()

		userPGRepo := NewUserPostgresRepo(mockDB, nil)

		userUUID := uuid.New().String()

		mockUser := &entity.User{
			ID:       userUUID,
			Email:    "test@example.com",
			Password: "hashed_password",
		}

		rows := pgxmock.NewRows([]string{"id", "email", "password"}).
			AddRow(
				userUUID,
				mockUser.Email,
				mockUser.Password,
			)

		expectedQuery := `INSERT INTO user (id,email,password) VALUES ($1,$2,$3) RETURNING id,email,password`

		mockDB.ExpectQuery(expectedQuery).WithArgs(
			mockUser.ID,
			mockUser.Email,
			mockUser.Password).WillReturnRows(rows)

		// Test the CreateUser function.
		createdUser, err := userPGRepo.CreateUser(context.Background(), mockUser)
		require.NoError(t, err)
		require.NotNil(t, createdUser)
	})
}

//func TestUserPostgresRepo_GetUserByID(t *testing.T) {
//
//	mockDB, err := pgxmock.NewPool()
//	require.NoError(t, err)
//
//	defer mockDB.Close()
//
//	userPGRepo := NewUserPostgresRepo(mockDB, nil)
//
//	userUUID := uuid.New().String()
//
//	mockUser := &entity.User{
//		ID:       userUUID,
//		Email:    "test@example.com",
//		Password: "hashed_password",
//	}
//
//	rows := pgxmock.NewRows([]string{"id", "email", "password"}).
//		AddRow(
//			userUUID,
//			mockUser.Email,
//			mockUser.Password,
//		)
//
//	expectedQuery := `SELECT id, email , password FROM "user" WHERE id = $1`
//
//	mockDB.ExpectQuery(expectedQuery).WithArgs(mockUser.ID).WillReturnRows(rows)
//
//	foundUser, err := userPGRepo.GetUserByID(context.Background(), mockUser.ID)
//	require.NoError(t, err)
//	require.NotNil(t, foundUser)
//	require.Equal(t, foundUser.ID, mockUser.ID)
//
//}
