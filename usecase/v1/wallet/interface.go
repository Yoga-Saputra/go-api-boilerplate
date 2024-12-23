package wallet

import (
	"github.com/Yoga-Saputra/go-boilerplate/internal/entity"
	"github.com/Yoga-Saputra/go-boilerplate/internal/repo"
)

type Repository interface {
	Transaction(txFunc func(interface{}) error) (err error)

	UpdateCredit(
		wallet *entity.Wallet,
		operator repo.WalletCreditOperator,
		amount float64,
		itx interface{},
		additionalUpdate ...repo.CustomUpdateStatements,
	) error

	Find(conds map[string]interface{}) (res entity.Wallet, rows int, err error)
	Finds(conds map[string]interface{}) (res []entity.Wallet, rows int, err error)
}
