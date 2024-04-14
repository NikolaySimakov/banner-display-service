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

func (br *BannerRepository) GetAdminBanners(
	ctx context.Context, 
	tagId int, 
	featureId int, 
	limit int64, 
	offset int64) ([]models.BannerResponse, error) {
	query := br.Builder.
        Select("id", "title", "text", "url", "created_at", "updated_at", 
        "last_version", "is_active", "tag_id", "feature_id").
        From("banners")

    if tagId != -1 {
        query = query.Where("? = ANY(tag_id)", tagId)
    }
    if featureId != -1 {
        query = query.Where("feature_id = ?", featureId)
    }
    if limit != -1 {
        query = query.Limit(uint64(limit))
    }
    if offset != -1 {
        query = query.Offset(uint64(offset))
    }

    sql, args, err := query.ToSql()
    if err != nil {
        return nil, err
    }

	rows, err := br.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("BannerRepository.GetAdminBanners - r.Pool.Query: %v", err)
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
			return nil, fmt.Errorf("BannerRepository.GetAdminBanners - rows.Scan: %v", err)
		}
		banners = append(banners, banner)
	}
	return banners, nil
}

func (br *BannerRepository) GetUserBanners(ctx context.Context, tagId int, featureId int, useLastRevision bool) ([]models.BannerResponse, error) {
	sql, args, _ := br.Builder.
		Select("id", "title", "text", "url", "created_at", "updated_at", 
		"last_version", "is_active", "tag_id", "feature_id").
		From("banners").
		Where("is_active = true AND ? = ANY(tag_id) AND feature_id = ? AND last_version = ?", tagId, featureId, useLastRevision).
		ToSql()

	rows, err := br.Pool.Query(ctx, sql, args...)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("BannerRepository.GetUserBanners - r.Pool.Query: %v", err)
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
			return nil, fmt.Errorf("BannerRepository.GetUserBanners - rows.Scan: %v", err)
		}
		banners = append(banners, banner)
	}
	return banners, nil
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