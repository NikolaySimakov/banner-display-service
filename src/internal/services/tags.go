package services

import (
	"banner-display-service/src/internal/repositories"
	"banner-display-service/src/internal/repositories/errs"
	"context"
	"errors"
	"fmt"
)

type TagService struct {
	tagRepo repositories.Tag
}

func NewTagService(tagRepo repositories.Tag) *TagService {
	return &TagService{
		tagRepo: tagRepo,
	}
}

func (ts *TagService) CreateTag(ctx context.Context, input TagInput) error {
	err := ts.tagRepo.CreateTag(ctx, input.Name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return nil
}

func (ts *TagService) DeleteTag(ctx context.Context, input TagInput) error {

	err := ts.tagRepo.DeleteTag(ctx, input.Name)

	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrNotFound
		}
		return err
	}

	return nil
}