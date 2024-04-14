package v1

import (
	"banner-display-service/src/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type featureRoutes struct {
	featureService services.Feature
	authService services.Auth
}

func newFeatureRoutes(g *echo.Group, featureService services.Feature, authService services.Auth) {
	r := featureRoutes{
		featureService: featureService,
		authService: authService,
	}
	g.POST("/", r.create)
	g.DELETE("/", r.delete)
}


type createFeatureInput struct {
	Name string `json:"slug" validate:"required,max=256"`
}

type createFeatureResponse struct {
	Name string `json:"name"`
}

func (f *featureRoutes) create(c echo.Context) error {

	// check admin token
	token := c.Request().Header.Get("token")
	if token == "" {
		newErrorResponse(c, http.StatusInternalServerError, ErrCannotParseToken.Error())
		return ErrCannotParseToken
	}

	userStatus, err := f.authService.TokenExist(c.Request().Context(), token)
	if err != nil || userStatus != "admin" {
		newErrorResponse(c, http.StatusInternalServerError, ErrInvalidAuthHeader.Error())
		return ErrInvalidAuthHeader
	}

	// create feature 
	var input createFeatureInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = f.featureService.CreateFeature(c.Request().Context(), services.FeatureInput{
		Name: input.Name,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusCreated, &createFeatureResponse{
		Name: input.Name,
	})
}


type deleteFeatureInput struct {
	Name string `json:"slug" validate:"required,max=256"`
}

func (f *featureRoutes) delete(c echo.Context) error {

	// check admin token
	token := c.Request().Header.Get("token")
	if token == "" {
		newErrorResponse(c, http.StatusInternalServerError, ErrCannotParseToken.Error())
		return ErrCannotParseToken
	}

	userStatus, err := f.authService.TokenExist(c.Request().Context(), token)
	if err != nil || userStatus != "admin" {
		newErrorResponse(c, http.StatusInternalServerError, ErrInvalidAuthHeader.Error())
		return ErrInvalidAuthHeader
	}
	
	// delete feature
	var input deleteFeatureInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = f.featureService.DeleteFeature(c.Request().Context(), services.FeatureInput{
		Name: input.Name,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.NoContent(204)
}