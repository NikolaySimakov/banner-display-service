package v1

import (
	"banner-display-service/src/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type tagRoutes struct {
	tagService services.Tag
	authService services.Auth
}

func newTagRoutes(g *echo.Group, tagService services.Tag, authService services.Auth) {
	r := tagRoutes{
		tagService: tagService,
		authService: authService,
	}
	g.POST("/", r.create)
	g.DELETE("/", r.delete)
}


type createTagInput struct {
	Name string `json:"slug" validate:"required,max=256"`
}

type createTagResponse struct {
	Name string `json:"name"`
}

func (t *tagRoutes) create(c echo.Context) error {

	// check admin token
	token := c.Request().Header.Get("token")
	if token == "" {
		newErrorResponse(c, http.StatusInternalServerError, ErrCannotParseToken.Error())
		return ErrCannotParseToken
	}

	userStatus, err := t.authService.TokenExist(c.Request().Context(), token)
	if err != nil || userStatus != "admin" {
		newErrorResponse(c, http.StatusInternalServerError, ErrInvalidAuthHeader.Error())
		return ErrInvalidAuthHeader
	}

	// create tag
	var input createTagInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = t.tagService.CreateTag(c.Request().Context(), services.TagInput{
		Name: input.Name,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.JSON(http.StatusCreated, &createTagResponse{
		Name: input.Name,
	})
}


type deleteTagInput struct {
	Name string `json:"slug" validate:"required,max=256"`
}

func (t *tagRoutes) delete(c echo.Context) error {

	// check admin token
	token := c.Request().Header.Get("token")
	if token == "" {
		newErrorResponse(c, http.StatusInternalServerError, ErrCannotParseToken.Error())
		return ErrCannotParseToken
	}

	userStatus, err := t.authService.TokenExist(c.Request().Context(), token)
	if err != nil || userStatus != "admin" {
		newErrorResponse(c, http.StatusInternalServerError, ErrInvalidAuthHeader.Error())
		return ErrInvalidAuthHeader
	}

	// delete tag
	var input deleteTagInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err = t.tagService.DeleteTag(c.Request().Context(), services.TagInput{
		Name: input.Name,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.NoContent(204)
}