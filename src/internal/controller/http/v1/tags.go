package v1

import (
	"banner-display-service/src/internal/services"

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


func (s *tagRoutes) create(c echo.Context) error {
	return nil
}

func (s *tagRoutes) delete(c echo.Context) error {
	return nil
}