package services

import (
	"banner-display-service/src/internal/repositories"
	"banner-display-service/src/pkg/secure"
	"context"
)

type AuthService struct {
	authRepo repositories.Auth
	secure   secure.APISecure
}

func NewAuthService(authRepo repositories.Auth, secure secure.APISecure) *AuthService {
	return &AuthService{
		authRepo: authRepo,
		secure:   secure,
	}
}

func (a *AuthService) TokenExist(ctx context.Context, token string) (string, error) {
	return a.authRepo.TokenExist(ctx, a.secure.Hash(token))
}

func (a *AuthService) GenerateToken(ctx context.Context, userStatus string) (int, string, error) {
	token := a.secure.GenerateKey()
	id, err := a.authRepo.WriteToken(ctx, a.secure.Hash(token), userStatus)
	return id, token, err
}