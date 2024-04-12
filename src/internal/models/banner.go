package models

type BannerModel struct {
	Id int
	Feature FeatureModel
	Tag []TagModel
	IsActive bool
}

// import (
// 	"database/sql"
// 	"time"
// )

// type Banner struct {
// 	Title     string `json:"title"`
// 	Text      string `json:"text"`
// 	Url       string `json:"url"`
// 	Active    bool   `json:"use_active-version"`
// 	CreatedAt time.Time
// 	UpdatedAt sql.NullTime
// }

// type UserBanner struct {
// 	Title  string `json:"title"`
// 	Text   string `json:"text"`
// 	Url    string `json:"url"`
// 	Active bool   `json:"use_active-version"`
// }