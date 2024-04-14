package db

import (
	"banner-display-service/src/pkg/postgres"
	"context"
	"fmt"
)

type AuthRepository struct {
	*postgres.Postgres
}

func NewAuthRepo(pg *postgres.Postgres) *AuthRepository {
	return &AuthRepository{pg}
}

func (a *AuthRepository) WriteToken(ctx context.Context, token string, userStatus string) (int, error) {

	fmt.Println(token, userStatus)

	sql, args, _ := a.Builder.
		Insert("api_keys").
		Columns("hash_key", "user_status").
		Values(token, userStatus).
		Suffix("RETURNING id").
		ToSql()

	var tokenID int
	err := a.Pool.QueryRow(ctx, sql, args...).Scan(&tokenID)
	if err != nil {
		return 0, fmt.Errorf("AuthRepository.WriteToken - u.Pool.QueryRow: %v", err)
	}
	return tokenID, nil
}

func (a *AuthRepository) TokenExist(ctx context.Context, token string) (string, error) {
	sql, args, _ := a.Builder.
		Select("user_status").
		From("api_keys").
		Where("hash_key = ?", token).
		ToSql()

	var userStatus string
	err := a.Pool.QueryRow(ctx, sql, args...).Scan(&userStatus)
	if err != nil {
		return "", fmt.Errorf("AuthRepository.TokenExist - u.Pool.QueryRow: %v", err)
	}
	return userStatus, nil
}