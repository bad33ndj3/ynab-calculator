package main

import (
	"flag"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"go.bmvs.io/ynab"
)

const (
	valuta = "€"
)

func main() {
	// create a fyne app
	a := app.New()
	w := a.NewWindow("Hello")

	// get flag of ynab oauth token
	ynabToken := flag.String("token", "", "YNAB OAuth token")
	flag.Parse()
	if ynabToken == nil {
		log.Panicln("Please provide a YNAB OAuth token")
	}

	// ynab api client
	c := ynab.NewClient(*ynabToken)

	// get budget tables
	bt := NewBudgetTableFetcher(c)
	objects, err := bt.Get()
	if err != nil {
		log.Panicln(err)
	}

	// render the table
	w.SetContent(container.New(layout.NewVBoxLayout(), objects...))
	w.ShowAndRun()
}
