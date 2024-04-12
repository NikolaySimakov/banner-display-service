package v1

import (
	"banner-display-service/src/internal/services"

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


func (f *featureRoutes) create(c echo.Context) error {
	return nil
}

func (f *featureRoutes) delete(c echo.Context) error {
	return nil
}