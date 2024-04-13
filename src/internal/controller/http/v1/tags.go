package v1

import (
	"banner-display-service/src/internal/services"
	"net/http"

	"github.com/labstack/echo"
)

type tagRoutes struct {
	tagService services.Tag
}

func newTagRoutes(g *echo.Group, tagService services.Tag) {
	r := tagRoutes{
		tagService: tagService,
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
	var input createTagInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := t.tagService.CreateTag(c.Request().Context(), services.TagInput{
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

	var input deleteTagInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := t.tagService.DeleteTag(c.Request().Context(), services.TagInput{
		Name: input.Name,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.NoContent(204)
}