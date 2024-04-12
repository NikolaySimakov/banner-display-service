package services

import (
	"banner-display-service/src/internal/repositories"
	"banner-display-service/src/internal/repositories/errs"
	"context"
	"errors"
	"fmt"
)

type FeatureService struct {
	featureRepo repositories.Feature
}

func NewFeatureService(featureRepo repositories.Feature) *FeatureService {
	return &FeatureService{
		featureRepo: featureRepo,
	}
}

func (fs *FeatureService) CreateFeature(ctx context.Context, input FeatureInput) error {
	err := fs.featureRepo.CreateFeature(ctx, input.name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

func (fs *FeatureService) DeleteFeature(ctx context.Context, input FeatureInput) error {
	
	err := fs.featureRepo.DeleteFeature(ctx, input.name)

	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrNotFound
		}
		return err
	}

	return nil
}