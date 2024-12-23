package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Yoga-Saputra/go-boilerplate/internal/entity"
	"gorm.io/gorm"
)

type (
	WalletCreditOperator string

	WalletRepoDB struct {
		db  *gorm.DB
		sql *sql.DB
	}
)

// Enum for wallet credit opperations
const (
	WalletAddCreditOp       WalletCreditOperator = "+"
	WalletSubstractCreditOp WalletCreditOperator = "-"
	WalletMultiplyCreditOp  WalletCreditOperator = "*"
	WalletDivideCreditOp    WalletCreditOperator = "/"
)

// NewWalletRepoDB create new DB repo for wallet entity.
func NewWalletRepoDB(db *gorm.DB, sql *sql.DB) *WalletRepoDB {
	if db != nil {
		return &WalletRepoDB{
			db:  db,
			sql: sql,
		}
	}

	return nil
}

// Validating operator contant enum
func (op WalletCreditOperator) validate() bool {
	valid := false

	switch op {
	case WalletAddCreditOp:
		valid = true

	case WalletSubstractCreditOp:
		valid = true

	case WalletMultiplyCreditOp:
		valid = true

	case WalletDivideCreditOp:
		valid = true
	}

	return valid
}

// Transaction repo method of wallet that approach DB process with transaction.
func (wrd *WalletRepoDB) Transaction(txFunc func(interface{}) error) (err error) {
	tx := wrd.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if err != nil {
			log.Printf("[DBTxSession] - Rollback with reason: %s", err.Error())
			tx.Rollback()
		} else {
			err = tx.Commit().Error
			if err != nil {
				log.Printf("[DBTxSession] - Commit error: %s", err.Error())
			}
		}
	}()

	err = txFunc(tx)
	return
}

// Find wallet records by given conditions and return wallet.
func (wrd *WalletRepoDB) Find(conds map[string]interface{}) (res entity.Wallet, rows int, err error) {
	// tx := wrd.db.Clauses(clause.Locking{Strength: "UPDATE"})

	tx := wrd.db.Find(&res, conds)
	rows = int(tx.RowsAffected)
	err = tx.Error

	return
}

// Finds wallet records by given conditions and return slice of wallets.
func (wrd *WalletRepoDB) Finds(conds map[string]interface{}) (res []entity.Wallet, rows int, err error) {
	tx := wrd.db.Find(&res, conds)
	rows = int(tx.RowsAffected)
	err = tx.Error

	return
}

func (wrd *WalletRepoDB) FindsWalletCommonLimit(conds map[string]interface{}) (res []entity.Wallet, err error) {
	var wSeamlessCheck entity.Wallet
	err = wrd.db.Order("id DESC").Limit(1).Find(&wSeamlessCheck, conds).Error
	if err != nil {
		return nil, err
	}

	lastSeenID := 0 // Initialize to 0 or a suitable starting point
	for {

		rows, err := wrd.db.Model(&entity.Wallet{}).Where("id > ?", lastSeenID).Select("id,member_id,amount,net_profit_loss,username,currency,is_new,is_locked,is_disabled,created_at,category,branch_id").
			Order("id ASC").
			Where(conds).
			Limit(300000).
			Rows()

		if err != nil {
			return nil, err
		}

		defer rows.Close()

		for rows.Next() {
			var wSeamless entity.Wallet

			if err = rows.Scan(&wSeamless.ID, &wSeamless.MemberID, &wSeamless.Amount, &wSeamless.NetProfitLoss, &wSeamless.Username, &wSeamless.Currency, &wSeamless.IsNew, &wSeamless.IsLocked, &wSeamless.IsDisabled, &wSeamless.CreatedAt, &wSeamless.Category, &wSeamless.BranchID); err != nil {
				return nil, err
			}
			res = append(res, wSeamless)
			lastSeenID = int(wSeamless.ID)

		}

		// Check if we got results
		if len(res) == 0 {
			break // Exit loop if no more results
		}

		if wSeamlessCheck.ID == uint64(lastSeenID) {
			break // Exit loop if no more results
		}
	}

	return res, nil
}

// Update wallet by given update statements and conditions.
func (wrd *WalletRepoDB) UpdateCredit(
	wallet *entity.Wallet,
	operator WalletCreditOperator,
	amount float64,
	itx interface{},
	additionalUpdate ...CustomUpdateStatements,
) error {
	// Validate argument
	switch {
	case amount < 0:
		return errors.New("amount cannot minus on UpdateCredit")
	case !operator.validate():
		return fmt.Errorf("mismatch operator type: %s", operator)
	case wallet == nil:
		return errors.New("argument wallet cannot nil pointer")
	}

	// Prepare update statements
	stats := map[string]interface{}{
		"amount": gorm.Expr("amount "+string(operator)+" ?", amount),
	}
	for _, au := range additionalUpdate {
		if au.UsingExpression {
			stats[au.Column] = gorm.Expr(au.Expr, au.Statement...)
		} else {
			stats[au.Column] = au.Statement
		}
	}

	// Prepare transaction session
	tx := wrd.db
	if itx != nil {
		itxdb, ok := itx.(*gorm.DB)
		if !ok {
			return errors.New("cannot assert argument itx to the *gorm.DB type")

		}

		tx = itxdb
	}

	// Update
	if err := tx.Model(&wallet).Updates(stats).Error; err != nil {
		return err
	}

	return nil
}
