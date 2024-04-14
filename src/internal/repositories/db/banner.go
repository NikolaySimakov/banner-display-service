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

func (br *BannerRepository) GetAllBanners(ctx context.Context) ([]models.BannerResponse, error) {
	sql, args, _ := br.Builder.
		Select("id", "title", "text", "url", "created_at", "updated_at", 
		"last_version", "is_active", "tag_id", "feature_id").
		From("banners").
		ToSql()

	rows, err := br.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("BannerRepository.GetAllBanners - r.Pool.Query: %v", err)
	}
	defer rows.Close()

	banners := make([]models.BannerResponse, 0, 1)
	for rows.Next() {
		banner := models.BannerResponse{}
		err = rows.Scan(
			&banner.Id, 
			&banner.Title,
			&banner.Text, 
			&banner.Url, 
			&banner.CreatedAt, 
			&banner.UpdatedAt, 
			&banner.LastVersion, 
			&banner.IsActive,
			&banner.TagId,
			&banner.FeatureId)
		if err != nil {
			return nil, fmt.Errorf("BannerRepository.GetAllBanners - rows.Scan: %v", err)
		}
		banners = append(banners, banner)
	}
	return banners, nil
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

func (br *BannerRepository) DeleteBanner(ctx context.Context, featureId int, tagId int) error {
	sql, args, _ := br.Builder.
		Delete("banners").
		Where("feature_id = ? AND ? = ANY(tag_id)", featureId, tagId).
		ToSql()

	commandTag, err := br.Pool.Exec(ctx, sql, args...)
	if err != nil {
			return fmt.Errorf("BannerRepository.DeleteBanner - br.Pool.Exec: %v", err)
	}

	if commandTag.RowsAffected() == 0 {
			return errs.ErrNotFound
	}

	return nil
}