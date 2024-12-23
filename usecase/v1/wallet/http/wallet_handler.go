package ucv1wallethttp

import (
	"time"

	"github.com/Yoga-Saputra/go-boilerplate/internal/entity/std"
	"github.com/labstack/echo/v4"
)

func (ds *domainService) Tetsing(c echo.Context) error {
	var apiResp *std.APIResponse

	now := time.Now()
	nowFormat := now.Format("2006-01-02 15:04:05")
	republish := map[string]interface{}{
		"last_update": nowFormat,
		"App":         "update wallet common",
	}

	apiResp = std.APIResponseSuccess(republish)
	return c.JSON(int(apiResp.StatusCode), apiResp.Body)
}
