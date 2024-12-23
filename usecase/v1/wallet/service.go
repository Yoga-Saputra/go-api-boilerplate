package wallet

import (
	"time"

	"github.com/Yoga-Saputra/go-boilerplate/internal/entity"
	"github.com/Yoga-Saputra/go-boilerplate/internal/repo"
	"github.com/Yoga-Saputra/go-boilerplate/pkg/grpcadp"
	"github.com/Yoga-Saputra/go-boilerplate/pkg/kemu"
)

// Service represent Republish Kafka services interface
type Service struct {
	repo    Repository
	kemu    *kemu.Mutex
	grpcAdp *grpcadp.WalletGrpcConn
}

// NewService creates new Republish Kafka services
func NewService(
	kemu *kemu.Mutex,
	r Repository,
	grpcAdp *grpcadp.WalletGrpcConn,
	callback ...func(s string),
) *Service {
	if len(callback) > 0 {
		callback[0]("Registering Wallet List Domain Entity...")
	}

	svc := &Service{
		repo:    r,
		kemu:    kemu,
		grpcAdp: grpcAdp,
	}

	// Set pointer value of export/public method
	newPublicAPI(svc)

	return svc
}

// AddWalletCreditById update wallet credit with given amount by given id
func (s *Service) UpdateWalletCommonCredit(
	logAmount, amount float64,
	now time.Time,
	wallet *entity.Wallet,
	operator repo.WalletCreditOperator,
	i interface{},
	additionalUpdateStmt []map[string]interface{},
) (err error) {
	// Update wallet credit

	customUpdate := []repo.CustomUpdateStatements{
		{
			UsingExpression: true,
			Column:          "net_profit_loss",
			Expr:            "net_profit_loss + ?",
			Statement:       []interface{}{logAmount},
		},
		{
			UsingExpression: false,
			Column:          "updated_at",
			Statement:       []interface{}{now},
		},
	}
	for _, m := range additionalUpdateStmt {
		for k, v := range m {
			customUpdate = append(customUpdate, repo.CustomUpdateStatements{
				UsingExpression: false,
				Column:          k,
				Statement:       []interface{}{v},
			})
		}
	}

	err = public.service.repo.UpdateCredit(
		wallet,
		operator,
		amount,
		i,
		customUpdate...,
	)
	return
}
