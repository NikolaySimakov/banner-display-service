package services

import (
	"banner-display-service/src/internal/models"
	"banner-display-service/src/internal/repositories"
	"banner-display-service/src/pkg/secure"
	"context"
)

type BannerInput struct {
	Id int
}

type Banner interface {
	GetAllBanners(ctx context.Context, userStatus string) ([]models.BannerResponse, error)
	GetUserBanner(ctx context.Context, input BannerInput) error
	CreateBanner(ctx context.Context, input *models.CreateBannerInput) error
	UpdateBanner(ctx context.Context, input BannerInput) error
	DeleteBanner(ctx context.Context, featureId int, tagId int) error
}

type TagInput struct {
	Name string
}

type Tag interface {
	CreateTag(ctx context.Context, input TagInput) error
	DeleteTag(ctx context.Context, input TagInput) error
}

type FeatureInput struct {
	Name string
}

type Feature interface {
	CreateFeature(ctx context.Context, input FeatureInput) error
	DeleteFeature(ctx context.Context, input FeatureInput) error
}

type Auth interface {
	TokenExist(ctx context.Context, token string) (string, error)
	GenerateToken(ctx context.Context, userStatus string) (int, string, error)
}

type Services struct {
	Banner
	Tag
	Feature
	Auth
}

type ServicesDependencies struct {
	Repos     *repositories.Repositories
	APISecure secure.APISecure
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Banner: NewBannerService(deps.Repos.Banner),
		Feature: NewFeatureService(deps.Repos.Feature),
		Tag: NewTagService(deps.Repos.Tag),
		Auth: NewAuthService(deps.Repos.Auth, deps.APISecure),
	}
}