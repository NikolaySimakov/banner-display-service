package services

import (
	"banner-display-service/src/internal/repositories"
	"context"
)

type GetBannerInput struct {
	Id int
}

type Banner interface {
	GetBanner(ctx context.Context, input GetBannerInput) error
}

type TagInput struct {
	name string
}

type Tag interface {
	CreateTag(ctx context.Context, input TagInput) error
	DeleteTag(ctx context.Context, input TagInput) error
}

type FeatureInput struct {
	name string
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