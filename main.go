package main

import (
	"flag"
	"fmt"
	"github.com/bad33ndj3/ynab/pkg/account"
	"github.com/leekchan/accounting"
	"go.bmvs.io/ynab"
)

const valuta = "€"

func main() {
	// get flag of ynab oauth token
	ynabToken := flag.String("token", "", "YNAB OAuth token")
	flag.Parse()
	if ynabToken == nil {
		fmt.Println("Please provide a YNAB OAuth token")
		return
	}

	// ynab api client
	c := ynab.NewClient(*ynabToken)
	budgets, err := c.Budget().GetBudgets()
	if err != nil {
		panic(err)
	}

	// print the total of each account
	for _, budget := range budgets {
		// if the budget name is T, skip it
		if budget.Name == "T" {
			continue
		}
		fmt.Println("Budget:", budget.Name)
		accounts, err := c.Account().GetAccounts(budget.ID)
		if err != nil {
			return
		}

		accountIterator := account.NewIterator(account.NewCollection(accounts))
		accountCalculator := account.NewCalculator(accountIterator)

		ac := accounting.Accounting{Symbol: valuta, Precision: 2}
		fmt.Printf("Total: %s\n", ac.FormatMoney(accountCalculator.Total()/1000))
		fmt.Printf("Total Stock: %s\n", ac.FormatMoney(accountCalculator.TotalStock()/1000))
		fmt.Printf("Total Crypto: %s\n", ac.FormatMoney(accountCalculator.TotalCrypto()/1000))
		fmt.Printf("Total Cash: %s\n", ac.FormatMoney(accountCalculator.TotalCash()/1000))
	}
}
