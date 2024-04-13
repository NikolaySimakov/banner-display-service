package services

import (
	"banner-display-service/src/internal/models"
	"banner-display-service/src/internal/repositories"
	"context"
)

type BannerInput struct {
	Id int
}

type Banner interface {
	GetAllBanners(ctx context.Context) error
	GetUserBanner(ctx context.Context, input BannerInput) error
	CreateBanner(ctx context.Context, input *models.CreateBannerInput) error
	UpdateBanner(ctx context.Context, input BannerInput) error
	DeleteBanner(ctx context.Context, input BannerInput) error
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

type Services struct {
	Banner
	Tag
	Feature
}

type ServicesDependencies struct {
	Repos     *repositories.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Banner: NewBannerService(deps.Repos.Banner),
		Feature: NewFeatureService(deps.Repos.Feature),
		Tag: NewTagService(deps.Repos.Tag),
	}
}