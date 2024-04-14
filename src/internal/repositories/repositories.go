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
	GetAdminBanners(ctx context.Context, tagId int, featureId int, limit int64, offset int64) ([]models.BannerResponse, error)
	GetUserBanners(ctx context.Context, tagId int, featureId int, useLastRevision bool) ([]models.BannerResponse, error)
	CreateBanner(ctx context.Context, banner *models.CreateBannerInput) error
	UpdateBanner(ctx context.Context, tag int, feature int) error
	DeleteBanner(ctx context.Context, featureId int, tagId int) error
}

type Auth interface {
	WriteToken(ctx context.Context, token string, userStatus string) (int, error)
	TokenExist(ctx context.Context, token string) (string, error)
}

type Repositories struct {
	Tag
	Feature
	Banner
	Auth
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Tag:       db.NewTagRepo(pg),
		Feature:    db.NewFeatureRepo(pg),
		Banner:    db.NewBannerRepo(pg),
		Auth:       db.NewAuthRepo(pg),
	}
}