package db

import (
	"banner-display-service/src/pkg/postgres"
	"context"
)

type BannerRepository struct {
	*postgres.Postgres
}

func NewBannerRepo(pg *postgres.Postgres) *BannerRepository {
	return &BannerRepository{pg}
}

func (br *BannerRepository) GetBanner(ctx context.Context, bannerId int) error {
	// TODO
	return nil
}

func (br *BannerRepository) CreateBanner(ctx context.Context, tag int, feature int) error {
	// TODO
	return nil
}

func (br *BannerRepository) UpdateBanner(ctx context.Context, tag int, feature int) error {
	// TODO
	return nil
}

func (br *BannerRepository) DeleteBanner(ctx context.Context) error {
	// TODO
	return nil
}