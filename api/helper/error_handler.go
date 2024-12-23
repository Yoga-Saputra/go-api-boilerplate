package helper

import (
	"errors"

	"github.com/Yoga-Saputra/go-boilerplate/internal/entity/std"
	"github.com/labstack/echo/v4"
)

// JsonErrorHandlerConfig defines the config for Json Error Handler helper
type JsonErrorHandlerConfig struct {
	// Err a error
	Err error

	// HTTPCode define http status code sended to response
	HTTPCode int
}

// DefaultJsonErrorHandlerConfig is the default config
var DefaultJsonErrorHandlerConfig = JsonErrorHandlerConfig{
	Err:      errors.New("opps, an error occured"),
	HTTPCode: int(std.StatusServerError),
}

// JsonErrorHandler return Json Error Handler helper.
// If config not given, will use the default
func (cfg JsonErrorHandlerConfig) JsonErrorHandler(err error, c echo.Context) {
	// Get config
	cfg = defaultJsonErrorHandlerConfig(cfg)

	// Default HTTP satatus code
	code := cfg.HTTPCode
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	// Set error
	if resp := std.APIResponseError(std.HTTPStatusCode(code), cfg.Err); resp != nil {
		c.JSON(int(resp.StatusCode), resp.Body)
	}
}

// Helper function to set default config
func defaultJsonErrorHandlerConfig(cfg JsonErrorHandlerConfig) JsonErrorHandlerConfig {
	// Overide default config
	c := cfg

	if c.Err == nil {
		c.Err = DefaultJsonErrorHandlerConfig.Err
	}
	if c.HTTPCode == 0 {
		c.HTTPCode = DefaultJsonErrorHandlerConfig.HTTPCode
	}

	return c
}
