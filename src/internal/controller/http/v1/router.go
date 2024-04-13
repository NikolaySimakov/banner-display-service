package v1

import (
	"banner-display-service/src/internal/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(handler *echo.Echo, services *services.Services) {
	// handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: `{"time":"${time_rfc3339_nano}", "method":"${method}","uri":"${uri}", "status":${status},"error":"${error}"}` + "\n",
	// 	Output: setLogsFile(),
	// }))
	handler.Use(middleware.Recover())

	handler.GET("/health", func(c echo.Context) error { return c.NoContent(200) })
	// handler.GET("/swagger/*", echoSwagger.WrapHandler)

	// authMiddleware := &AuthMiddleware{services.Auth}
	// v1 := handler.Group("/api/v1", authMiddleware.UserIdentity)
	v1 := handler.Group("/api/v1")
	{
		newTagRoutes(v1.Group("/tags"), services.Tag)
		newFeatureRoutes(v1.Group("/features"), services.Feature)
		newBannerRoutes(v1.Group("/banners"), services.Banner)
	}
}