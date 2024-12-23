package ucv1wallethttp

import (
	"errors"

	"github.com/Yoga-Saputra/go-boilerplate/internal/entity/std"
	"github.com/labstack/echo/v4"
)

// ValReqP__WalletCreateNew validate request params/payloads
func (ds *domainService) StructValidator(c echo.Context, p interface{}) *std.APIResponse {
	// Parsing params/payload to struct
	if err := c.Bind(p); err != nil {
		return std.APIResponseError(std.StatusBadRequest, errors.New("failed to parsing request params/payloads"))
	}

	// Validate the params/payloads
	if err := c.Validate(p); err != nil {
		return std.APIResponseError(std.StatusBadRequest, err)
	}

	return nil
}
