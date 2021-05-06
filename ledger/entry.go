package ledger

import (
	"errors"

	"github.com/Rhymond/go-money"
)

const (
	debit  = "debit"
	credit = "credit"
)

// Operation represents valid accounting operations "credit" and "debit"
type Operation struct {
	name string
}

var operations = map[string]*Operation{
	debit:  {debit},
	credit: {credit},
}

func getOperation(name string) *Operation {
	return operations[name]
}

// Entry represents a value entry (money) on a given Account
type Entry struct {
	*Account
	*money.Money
	*Operation
}

// NewEntry returns a new Entry representing a value operation on a given Account
func (a *Account) NewEntry(operation string, amount int64, code string) (*Entry, error) {
	o := getOperation(operation)
	if o == nil {
		return nil, errors.New("operation must be either 'credit' or 'debit'")
	}
	if amount < 0 {
		return nil, errors.New("amount must be greater than 0")
	}
	c := money.GetCurrency(code)
	if c == nil {
		return nil, errors.New("Unknow currency code. Sure it's ISO 4217?")
	}
	m := money.New(amount, c.Code)

	return &Entry{a, m, o}, nil
}
