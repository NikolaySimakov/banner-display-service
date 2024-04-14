package v1

import (
	"banner-display-service/src/internal/models"
	"banner-display-service/src/internal/services"
	"net/http"

	"github.com/labstack/echo"
)

type bannerRoutes struct {
	bannerService services.Banner
	authService services.Auth
}

func newBannerRoutes(g *echo.Group, bannerService services.Banner, authService services.Auth) {
	r := bannerRoutes{
		bannerService: bannerService,
		authService: authService,
	}
	g.GET("/", r.readAll)
	g.GET("/:id", r.read)
	g.POST("/", r.create)
	g.DELETE("/", r.delete)
}


func (b *bannerRoutes) readAll(c echo.Context) error {

	// token := c.QueryParam("token")
	token := c.Request().Header.Get("token")
	if token == "" {
		newErrorResponse(c, http.StatusInternalServerError, ErrCannotParseToken.Error())
		return ErrCannotParseToken
	}

	userStatus, err := b.authService.TokenExist(c.Request().Context(), token)
	if err != nil || !(userStatus == "admin" || userStatus == "user") {
		newErrorResponse(c, http.StatusInternalServerError, ErrInvalidAuthHeader.Error())
		return ErrInvalidAuthHeader
	}

	banners, err := b.bannerService.GetAllBanners(c.Request().Context(), userStatus)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, banners)
}

func (b *bannerRoutes) read(c echo.Context) error {
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

	// check admin token
	token := c.Request().Header.Get("token")
	if token == "" {
		newErrorResponse(c, http.StatusInternalServerError, ErrCannotParseToken.Error())
		return ErrCannotParseToken
	}

	userStatus, err := b.authService.TokenExist(c.Request().Context(), token)
	if err != nil || userStatus != "admin" {
		newErrorResponse(c, http.StatusInternalServerError, ErrInvalidAuthHeader.Error())
		return ErrInvalidAuthHeader
	}

	// create banner
	var input createBannerInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = b.bannerService.CreateBanner(c.Request().Context(), &models.CreateBannerInput{
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


type deleteBannerInput struct {
	FeatureId int `json:"feature_id"`
	TagId int `json:"tag_id"`
}

func (b *bannerRoutes) delete(c echo.Context) error {

	// check admin token
	token := c.Request().Header.Get("token")
	if token == "" {
		newErrorResponse(c, http.StatusInternalServerError, ErrCannotParseToken.Error())
		return ErrCannotParseToken
	}

	userStatus, err := b.authService.TokenExist(c.Request().Context(), token)
	if err != nil || userStatus != "admin" {
		newErrorResponse(c, http.StatusInternalServerError, ErrInvalidAuthHeader.Error())
		return ErrInvalidAuthHeader
	}

	// delete banner
	var input deleteBannerInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = b.bannerService.DeleteBanner(c.Request().Context(), input.FeatureId, input.TagId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.NoContent(204)
}