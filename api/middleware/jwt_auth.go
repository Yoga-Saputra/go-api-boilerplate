package middleware

import (
	"errors"

	"github.com/Yoga-Saputra/go-boilerplate/config"
	"github.com/Yoga-Saputra/go-boilerplate/internal/entity/std"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	mdlr "github.com/labstack/echo/v4/middleware"
)

func JWTAuth() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		JWTValidateToken(),
		// JWTVerifySecretKey,
	}
}

func JWTValidateToken() func(next echo.HandlerFunc) echo.HandlerFunc {
	return mdlr.JWTWithConfig(mdlr.JWTConfig{
		SigningMethod: "RS256",
		SigningKey:    config.Of.App.GetPublicKey(),
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			apiResp := std.APIResponseError(std.StatusForbidden, errors.New("unauthorized"))
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		},
	})
}

func JWTVerifySecretKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		usr := c.Get("user").(*jwt.Token)
		claims := usr.Claims.(jwt.MapClaims)
		seck := claims["sec"].(string)

		if config.Of.App.GetSecretKey() != seck {
			err := errors.New("your token has been expired")
			apiResp := std.APIResponseError(std.StatusForbidden, err)
			return c.JSON(int(apiResp.StatusCode), apiResp.Body)
		}

		return next(c)
	}
}
