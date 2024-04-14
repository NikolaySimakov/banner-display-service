package v1

import (
	"banner-display-service/src/internal/models"
	"banner-display-service/src/internal/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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
	g.GET("/banner", r.getAdminBanner)
	g.GET("/user_banner", r.getUserBanner)
	g.POST("/", r.create)
	g.DELETE("/", r.delete)
}


func (b *bannerRoutes) getAdminBanner(c echo.Context) error {

	// check user & admin token
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

	tagId := c.QueryParam("tag_id")
	featureId := c.QueryParam("feature_id")
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	tagIdInt, tagErr := strconv.Atoi(tagId)
	featureIdInt, featureErr := strconv.Atoi(featureId)
	limitInt, limitErr := strconv.ParseInt(limit, 10, 64)
	offsetInt, offsetErr := strconv.ParseInt(offset, 10, 64)
	if tagErr != nil {
		tagIdInt = -1
	}
	if featureErr != nil {
		featureIdInt = -1
	}
	if limitErr != nil {
		limitInt = -1
	}
	if offsetErr != nil {
		offsetInt = -1
	}

	banners, err := b.bannerService.GetAdminBanners(c.Request().Context(), tagIdInt, featureIdInt, limitInt, offsetInt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, banners)
}

func (b *bannerRoutes) getUserBanner(c echo.Context) error {

	// check user & admin token
	token := c.Request().Header.Get("token")
	if token == "" {
		newErrorResponse(c, http.StatusInternalServerError, ErrCannotParseToken.Error())
		return ErrCannotParseToken
	}

	userStatus, err := b.authService.TokenExist(c.Request().Context(), token)
	if err != nil || userStatus != "user" {
		newErrorResponse(c, http.StatusInternalServerError, ErrInvalidAuthHeader.Error())
		return ErrInvalidAuthHeader
	}

	tagId := c.QueryParam("tag_id")
	featureId := c.QueryParam("feature_id")
	useLastRevision := c.QueryParam("use_last_revision")

	tagIdInt, tagErr := strconv.Atoi(tagId)
	featureIdInt, featureErr := strconv.Atoi(featureId)
	useLastRevisionBool, useLastRevisionErr := strconv.ParseBool(useLastRevision)
	if tagErr != nil || featureErr != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	if useLastRevisionErr != nil {
		useLastRevisionBool = false
	}

	fmt.Println(tagIdInt, featureIdInt, useLastRevisionBool)

	banners, err := b.bannerService.GetUserBanners(c.Request().Context(), tagIdInt, featureIdInt, useLastRevisionBool)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusOK, banners)
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