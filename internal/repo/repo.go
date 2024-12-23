package repo

import (
	"fmt"

	"github.com/sutejoramadhan/gormjqdt"
	"gorm.io/gorm"
)

type (
	// RepoDatatableResponse defines datatable response from gormjqdt package
	RepoDatatableResponse gormjqdt.Response

	AdditionalConditions struct {
		Qry interface{}
		Arg []interface{}

		GroupingOr []AdditionalConditions
	}

	// CustomUpdateStatements hold data structtur for updating record using custom statement.
	CustomUpdateStatements struct {
		UsingExpression bool
		Column          string
		Expr            string
		Statement       []interface{}
	}
)

// DB Scope for limit query.
func limitScope(limit, offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

// DB Scope for filter by branchs query.
func filterBranchScope(branchs []int, columnName ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var qry interface{} = "branch_id in (?)"
		if len(columnName) > 0 {
			qry = fmt.Sprintf("%s in (?)", columnName[0])
		}

		return db.Where(qry, branchs)
	}
}

// DB Scope for filter by currencies query.
func filterCurrencyScope(currencies []string, columnName ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var qry interface{} = "currency in (?)"
		if len(columnName) > 0 {
			qry = fmt.Sprintf("%s in (?)", columnName[0])
		}

		return db.Where(qry, currencies)
	}
}

// DB Scope for filter by custom conditions query.
func filterCustomScope(conds map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(conds)
	}

}

// DB Scope for filter by advance custom conditions query.
func filterAdvanceCustomScope(
	conds []AdditionalConditions,
	usingOr bool,
	isGroupConds ...bool,
) func(db *gorm.DB) *gorm.DB {
	forGroupingConds := false
	if len(isGroupConds) > 0 {
		forGroupingConds = isGroupConds[0]
	}

	return func(db *gorm.DB) *gorm.DB {
		tx := db
		if forGroupingConds {
			tx = tx.Session(&gorm.Session{NewDB: true})
		}

		for _, v := range conds {
			if usingOr {
				tx = tx.Or(v.Qry, v.Arg...)
			} else {
				tx = tx.Where(v.Qry, v.Arg...)
			}
		}

		// If this scope using for grouped conditions return grouped query
		// otherwise return chained query
		if forGroupingConds {
			return db.Where(tx)
		} else {
			return tx
		}
	}
}
