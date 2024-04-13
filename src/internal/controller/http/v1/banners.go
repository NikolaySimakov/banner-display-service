package v1

import (
	"banner-display-service/src/internal/models"
	"banner-display-service/src/internal/services"
	"net/http"

	"github.com/labstack/echo"
)

type bannerRoutes struct {
	bannerService services.Banner
}

func newBannerRoutes(g *echo.Group, bannerService services.Banner) {
	r := bannerRoutes{
		bannerService: bannerService,
	}
	g.GET("/", r.readAll)
	g.POST("/", r.create)
	g.DELETE("/", r.delete)
}


func (b *bannerRoutes) readAll(c echo.Context) error {
	return nil
}

type createBannerInput struct {
	Title string `json:"title"`
	Text string `json:"text"`
	Url string `json:"url"`
	FeatureId int `json:"feature"`
	TagId []int `json:"tag"`
	IsActive bool `json:"active"`
}

func (b *bannerRoutes) create(c echo.Context) error {
	var input createBannerInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := b.bannerService.CreateBanner(c.Request().Context(), &models.CreateBannerInput{
		Title: input.Title,
		Text: input.Text,
		Url: input.Url,
		FeatureId: input.FeatureId,
		TagId: input.TagId,
		IsActive: input.IsActive,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return nil
}

func (b *bannerRoutes) delete(c echo.Context) error {
	return nil
}