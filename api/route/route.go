package route

import (
	"github.com/labstack/echo/v4"
)

// Register registers all defined route
func RegisterGroup(e *echo.Echo) map[string]*echo.Group {
	return map[string]*echo.Group{
		"v1": v1(e),
	}
}
