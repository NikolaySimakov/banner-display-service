package models

type CreateBannerInput struct {
	// Id int
	Title string `json:"title"`
	Text string `json:"text"`
	Url string `json:"url"`
	FeatureId int `json:"feature"`
	TagId []int `json:"tag"`
	IsActive bool `json:"active"`
}