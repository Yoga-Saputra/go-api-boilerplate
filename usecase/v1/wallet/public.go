package wallet

import (
	"errors"

	"github.com/Yoga-Saputra/go-boilerplate/internal/entity"
)

// Defines struct that will be exported
type publicAPI struct {
	service *Service
}

var public *publicAPI

func newPublicAPI(svc *Service) {
	public = &publicAPI{
		service: svc,
	}
}

type TopUpSyncPayload struct {
	PID          string
	ProviderCode string
	RefID        string
	ServiceID    uint32
	Amount       float64
	ProcessedBy  string
	Category     entity.WalletCategory
}

// Validating local variable.
func validatePointer() error {
	if public == nil {
		return errors.New("package cannot be accessed, lease implement the usecase first")
	}

	return nil
}
