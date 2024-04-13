package models

import (
	"database/sql"
	"time"
)

type CreateBannerInput struct {
	// Id int
	Title string `json:"title"`
	Text string `json:"text"`
	Url string `json:"url"`
	FeatureId int `json:"feature_id"`
	TagId []int `json:"tag_id"`
	IsActive bool `json:"is_active"`
}

type BannerResponse struct {
	Id            int             `db:"id"`
	Title         string          `db:"title"`
	Text          string          `db:"text"`
	Url           string          `db:"url"`
	CreatedAt     time.Time       `db:"created_at"`
	UpdatedAt     sql.NullTime    `db:"updated_at"`
	LastVersion   bool            `db:"last_version"`
	IsActive      bool            `db:"is_active"`
	TagId         []int           `db:"tag_id"`
	FeatureId     int             `db:"feature_id"`
}