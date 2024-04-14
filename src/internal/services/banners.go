package services

import (
	"banner-display-service/src/internal/models"
	"banner-display-service/src/internal/repositories"
	"banner-display-service/src/internal/repositories/errs"
	"context"
	"errors"
)

type BannerService struct {
	bannerRepo repositories.Banner
}

func NewBannerService(bannerRepo repositories.Banner) *BannerService {
	return &BannerService{
		bannerRepo: bannerRepo,
	}
}

func (bs *BannerService) GetAdminBanners(ctx context.Context, tagId int, featureId int, limit int64, offset int64) ([]models.BannerResponse, error) {
	banners, err := bs.bannerRepo.GetAdminBanners(ctx, tagId, featureId, limit, offset)

	if err != nil {
		return nil, err
	}

	return banners, nil
}

func (bs *BannerService) GetUserBanners(ctx context.Context, tagId int, featureId int, useLastRevision bool) ([]models.BannerResponse, error) {
	banners, err := bs.bannerRepo.GetUserBanners(ctx, tagId, featureId, useLastRevision)

	if err != nil {
		return nil, err
	}

	return banners, nil
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

func (bs *BannerService) DeleteBanner(ctx context.Context, featureId int, tagId int) error {
	err := bs.bannerRepo.DeleteBanner(ctx, featureId, tagId)

	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrNotFound
		}
		return err
	}

	return nil
}