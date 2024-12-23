package route

import (
	"github.com/Yoga-Saputra/go-boilerplate/api/middleware"
	"github.com/labstack/echo/v4"
)

func v1(e *echo.Echo) *echo.Group {
	return e.Group("/v1", middleware.JWTAuth()...)
}
