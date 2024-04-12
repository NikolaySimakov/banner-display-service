package services

import (
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

func (bs *BannerService) GetBanner(ctx context.Context, input GetBannerInput) error {
	return nil
}