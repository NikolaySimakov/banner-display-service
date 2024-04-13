package repositories

import (
	"banner-display-service/src/internal/models"
	"banner-display-service/src/internal/repositories/db"
	"banner-display-service/src/pkg/postgres"
	"context"
)

type Tag interface {
	CreateTag(ctx context.Context, name string) error
	DeleteTag(ctx context.Context, name string) error
}

type Feature interface {
	CreateFeature(ctx context.Context, name string) error
	DeleteFeature(ctx context.Context, name string) error
}

type Banner interface {
	GetAllBanners(ctx context.Context) error
	GetUserBanner(ctx context.Context, bannerId int) error
	CreateBanner(ctx context.Context, banner *models.CreateBannerInput) error
	UpdateBanner(ctx context.Context, tag int, feature int) error
	DeleteBanner(ctx context.Context) error
}

type Repositories struct {
	Tag
	Feature
	Banner
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Tag:       db.NewTagRepo(pg),
		Feature:    db.NewFeatureRepo(pg),
		Banner:    db.NewBannerRepo(pg),
	}
}