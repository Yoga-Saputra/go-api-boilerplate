package ucv1wallethttp

import (
	"github.com/Yoga-Saputra/go-boilerplate/pkg/kemu"
	"github.com/Yoga-Saputra/go-boilerplate/usecase/v1/wallet"

	"github.com/labstack/echo/v4"
)

type domainService struct {
	s    wallet.Service
	kemu *kemu.Mutex
}

func RegisterRoute(v1 *echo.Group, s wallet.Service, k *kemu.Mutex) {
	// Setup domain service
	ds := &domainService{
		s:    s,
		kemu: k,
	}

	// Create root wallet group
	wg := v1.Group("/wallet") // <- Route group (and also prefix) "wallet"

	wg.POST("/testing", ds.Tetsing) // <- wallet testing

}
