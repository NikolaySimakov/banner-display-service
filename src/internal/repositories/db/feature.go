package db

import (
	"banner-display-service/src/internal/repositories/errs"
	"banner-display-service/src/pkg/postgres"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
)

type FeatureRepository struct {
	*postgres.Postgres
}

func NewFeatureRepo(pg *postgres.Postgres) *FeatureRepository {
	return &FeatureRepository{pg}
}

func (fr *FeatureRepository) CreateFeature(ctx context.Context, name string) error {
	sql, args, _ := fr.Builder.
		Insert("features").
		Columns("name").
		Values(name).
		ToSql()

	err := fr.Pool.QueryRow(ctx, sql, args...).Scan()
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

		return fmt.Errorf("SegmentRepo.CreateSegment - s.Pool.QueryRow: %v", err)
	}
	return nil
}

func (fr *FeatureRepository) DeleteFeature(ctx context.Context, name string) error {
	sql, args, _ := fr.Builder.
		Delete("features").
		Where("name = ?", name).
		Suffix("RETURNING name").
		ToSql()

	err := fr.Pool.QueryRow(ctx, sql, args...).Scan(&name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errs.ErrNotFound
		}
		return fmt.Errorf("SegmentRepo.DeleteSegment - s.Pool.QueryRow: %v", err)
	}
	return nil
}