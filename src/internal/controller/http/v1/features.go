package v1

import (
	"banner-display-service/src/internal/services"
	"net/http"

	"github.com/labstack/echo"
)

type featureRoutes struct {
	featureService services.Feature
}

func newFeatureRoutes(g *echo.Group, featureService services.Feature) {
	r := featureRoutes{
		featureService: featureService,
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
	var input createFeatureInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := f.featureService.CreateFeature(c.Request().Context(), services.FeatureInput{
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
	
	var input deleteFeatureInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	err := f.featureService.DeleteFeature(c.Request().Context(), services.FeatureInput{
		Name: input.Name,
	})

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}

	return c.NoContent(204)
}