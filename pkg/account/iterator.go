package account

import (
	"go.bmvs.io/ynab/api/account"
)

type Iterator interface {
	Add(a *account.Account)
	Next() *account.Account
	HasNext() bool
	Reset()
	Index() int
	Count() int
	All() []*account.Account
	Get(index int) *account.Account
	GetByID(id string) *account.Account
	GetByName(name string) *account.Account
	GetByType(t account.Type) *account.Account
	NewIterator() Iterator
}

// NewIterator creates a new iterator for the account
func NewIterator(a *Collection) Iterator {
	return &accountIterator{
		account: a,
		index:   -1,
	}
}

// NewCollection creates a new collection of accounts
func NewCollection(accounts []*account.Account) *Collection {
	return &Collection{
		accounts: accounts,
	}
}

// Collection is a collection of account.Account accounts
type Collection struct {
	accounts []*account.Account
}

// accountIterator iterates over the accounts
type accountIterator struct {
	account *Collection
	index   int
}

func (i *accountIterator) NewIterator() Iterator {
	return &accountIterator{
		account: i.account,
		index:   -1,
	}
}

// Add adds the given account to the iterator
func (i *accountIterator) Add(a *account.Account) {
	i.account.accounts = append(i.account.accounts, a)
}

// Next returns the next account
func (i *accountIterator) Next() *account.Account {
	if !i.HasNext() {
		return nil
	}
	i.index++
	return i.account.accounts[i.index]
}

// HasNext returns true if there is another account
func (i *accountIterator) HasNext() bool {
	return i.index < len(i.account.accounts)-1
}

// Reset resets the iterator
func (i *accountIterator) Reset() {
	i.index = -1
}

// Index returns the current index
func (i *accountIterator) Index() int {
	return i.index
}

// Count returns the number of accounts
func (i *accountIterator) Count() int {
	return len(i.account.accounts)
}

// All returns all accounts
func (i *accountIterator) All() []*account.Account {
	return i.account.accounts
}

// Get returns the account at the given index
func (i *accountIterator) Get(index int) *account.Account {
	return i.account.accounts[index]
}

// GetByID returns the account with the given ID
func (i *accountIterator) GetByID(id string) *account.Account {
	for _, a := range i.account.accounts {
		if a.ID == id {
			return a
		}
	}
	return nil
}

// GetByName returns the account with the given name
func (i *accountIterator) GetByName(name string) *account.Account {
	for _, a := range i.account.accounts {
		if a.Name == name {
			return a
		}
	}
	return nil
}

// GetByType returns the account with the given type
func (i *accountIterator) GetByType(t account.Type) *account.Account {
	for _, a := range i.account.accounts {
		if a.Type == t {
			return a
		}
	}
	return nil
}
