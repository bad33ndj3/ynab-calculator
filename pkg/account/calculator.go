package account

import "strings"

// Calculator is a calculator interface.
type Calculator interface {
	// Total returns the total of the account.
	Total() float64
	// TotalCrypto returns the total of the account in crypto.
	TotalCrypto() float64
	// TotalStock returns the total of the account in stock.
	TotalStock() float64
	// TotalCash returns the total of the account in cash.
	TotalCash() float64
}

// NewCalculator returns a new Calculator. Tst
func NewCalculator(acc Iterator) Calculator {
	return &calculator{acc: acc}
}

// calculator implements Calculator.
type calculator struct {
	acc Iterator
}

func (c *calculator) Total() float64 {
	// return the total of balances
	var total float64
	defer c.acc.Reset()
	for c.acc.HasNext() {
		acc := c.acc.Next()
		total += c.int64toFloat64(acc.Balance)
	}
	return total
}

// TotalStock returns the total of the stock.
func (c *calculator) TotalStock() float64 {
	// return the total of balances
	var total float64
	defer c.acc.Reset()
	for c.acc.HasNext() {
		acc := c.acc.Next()
		// check if the account name contains "stock" in lowercase
		if strings.Contains(strings.ToLower(acc.Name), "stock") {
			total += c.int64toFloat64(acc.Balance)
		}
	}
	return total
}

// TotalCrypto returns the total of the crypto.
func (c *calculator) TotalCrypto() float64 {
	// return the total of balances
	var total float64
	defer c.acc.Reset()
	for c.acc.HasNext() {
		acc := c.acc.Next()
		// check if the account name contains "crypto" in lowercase
		if strings.Contains(strings.ToLower(acc.Name), "crypto") {
			total += c.int64toFloat64(acc.Balance)
		}
	}
	return total
}

// TotalCash returns the total of the cash.
func (c *calculator) TotalCash() float64 {
	// return the total of balances
	var total float64
	defer c.acc.Reset()
	for c.acc.HasNext() {
		acc := c.acc.Next()
		total += c.int64toFloat64(acc.Balance)
	}
	// get a var with the total of crypto on a new Iterator
	crypto := c.newCalculator(c.acc.NewIterator()).TotalCrypto()
	stock := c.newCalculator(c.acc.NewIterator()).TotalStock()

	return total - crypto - stock
}

// newCalculator returns a new Calculator.
func (c *calculator) newCalculator(acc Iterator) Calculator {
	return &calculator{acc: acc}
}

// int64toFloat64 converts an int64 to a float64.
func (c *calculator) int64toFloat64(i int64) float64 {
	return float64(i)
}
