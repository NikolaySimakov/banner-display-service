package services

import (
	"banner-display-service/src/internal/models"
	"banner-display-service/src/internal/repositories"
	"context"
)

type BannerService struct {
	bannerRepo repositories.Banner
}

func NewBannerService(bannerRepo repositories.Banner) *BannerService {
	return &BannerService{
		bannerRepo: bannerRepo,
	}
}

func (bs *BannerService) GetAllBanners(ctx context.Context) error {
	return nil
}

func (bs *BannerService) GetUserBanner(ctx context.Context, input BannerInput) error {
	return nil
}

func (bs *BannerService) CreateBanner(ctx context.Context, input *models.CreateBannerInput) error {
	err := bs.bannerRepo.CreateBanner(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

func (bs *BannerService) UpdateBanner(ctx context.Context, input BannerInput) error {
	return nil
}

func (bs *BannerService) DeleteBanner(ctx context.Context, input BannerInput) error {
	return nil
}