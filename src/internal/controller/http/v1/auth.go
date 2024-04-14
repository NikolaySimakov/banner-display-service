package v1

import (
	"banner-display-service/src/internal/services"
	"net/http"

	"github.com/labstack/echo"
)

type authRoutes struct {
	authService services.Auth
}

func newAuthRoutes(g *echo.Group, authService services.Auth) {
	r := authRoutes{
		authService: authService,
	}
	g.POST("/", r.create)
}

type createUserInput struct {
	UserStatus string `json:"user_status"`
}

type createUserResponse struct {
	UserId int `json:"user_id"`
	UserToken string `json:"user_token"`
	UserStatus string `json:"user_status"`
}

func (a *authRoutes) create(c echo.Context) error {
	var input createUserInput

	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return err
	}

	if err := c.Validate(input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return err
	}

	userId, userToken, err := a.authService.GenerateToken(c.Request().Context(), input.UserStatus)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return err
	}
	
	return c.JSON(http.StatusOK, &createUserResponse{
		UserId: userId,
		UserToken: userToken,
		UserStatus: input.UserStatus,
	})
}