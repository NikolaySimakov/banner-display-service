package db

import (
	"banner-display-service/src/internal/models"
	"banner-display-service/src/internal/repositories/errs"
	"banner-display-service/src/pkg/postgres"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
)

type BannerRepository struct {
	*postgres.Postgres
}

func NewBannerRepo(pg *postgres.Postgres) *BannerRepository {
	return &BannerRepository{pg}
}

func (br *BannerRepository) GetAllBanners(ctx context.Context) error {
	return nil
}

func (br *BannerRepository) GetUserBanner(ctx context.Context, bannerId int) error {
	// TODO
	return nil
}

func (br *BannerRepository) CreateBanner(ctx context.Context, banner *models.CreateBannerInput) error {

    sql, args, err := br.Builder.
        Insert("banners").
        Columns("title", "text", "url", "feature_id", "tag_id", "is_active").
        Values(banner.Title, banner.Text, banner.Url, banner.FeatureId, banner.TagId, banner.IsActive).
        ToSql()

    if err != nil {
        return fmt.Errorf("BannerRepository.CreateBanner - failed to build SQL query: %v", err)
    }

		// Execute the query
    _, err = br.Pool.Exec(ctx, sql, args...)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil
        }

        var pgErr *pgconn.PgError
        if ok := errors.As(err, &pgErr); ok {
            if pgErr.Code == "23505" {
                return errs.ErrAlreadyExists
            }
        }

        return fmt.Errorf("BannerRepository.CreateBanner - failed to execute SQL query: %v", err)
    }

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