package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/bad33ndj3/ynab-calculator/pkg/account"
	"github.com/leekchan/accounting"
	"go.bmvs.io/ynab"
)

const (
	toCents       = 1000
	coinPrecision = 2
)

// BudgetTableFetcher is an interface for fetching budgets.
type BudgetTableFetcher interface {
	Get() ([]fyne.CanvasObject, error)
}

// budgetTableFetcher is a struct that holds the budget tables.
type budgetTableFetcher struct {
	ynab ynab.ClientServicer
}

// NewBudgetTableFetcher returns a new budgetTableFetcher.
func NewBudgetTableFetcher(ynabClient ynab.ClientServicer) BudgetTableFetcher {
	return &budgetTableFetcher{ynab: ynabClient}
}

// Get returns the budget tables.
func (c *budgetTableFetcher) Get() ([]fyne.CanvasObject, error) {
	budgets, err := c.ynab.Budget().GetBudgets()
	if err != nil {
		return nil, err
	}

	var objects []fyne.CanvasObject

	// print the total of each account
	for _, budget := range budgets {
		// if the budget name is T, skip it
		if budget.Name == "T" {
			continue
		}

		// get the accounts
		accounts, err := c.ynab.Account().GetAccounts(budget.ID)
		if err != nil {
			return nil, err
		}

		// get the transactions
		accountIterator := account.NewIterator(account.NewCollection(accounts))
		accountCalculator := account.NewCalculator(accountIterator)

		information := make([][2]string, 0)

		// get the total of each account
		ac := accounting.Accounting{Symbol: valuta, Precision: coinPrecision}
		information = append(
			information,
			[2]string{"Total ", ac.FormatMoney(accountCalculator.Total() / toCents)},
			[2]string{"Total Stock ", ac.FormatMoney(accountCalculator.TotalStock() / toCents)},
			[2]string{"Total Crypto ", ac.FormatMoney(accountCalculator.TotalCrypto() / toCents)},
			[2]string{"Total Cash ", ac.FormatMoney(accountCalculator.TotalCash() / toCents)},
		)

		// add the total of each account to a table
		objects = append(objects,
			widget.NewLabel(fmt.Sprintf("Budget: %s", budget.Name)),
			widget.NewTable(func() (int, int) {
				if len(information) == 0 {
					return 0, 0
				}
				return len(information), len(information[0])
			}, func() fyne.CanvasObject {
				return widget.NewLabel("")
			}, func(i widget.TableCellID, o fyne.CanvasObject) {
				o.(*widget.Label).SetText(information[i.Row][i.Col])
			}),
			layout.NewSpacer(),
		)
	}

	return objects, nil
}
