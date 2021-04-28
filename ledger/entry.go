package ledger

import (
	"fmt"
)

type Entry struct {
	*Account
	Debit    int
	Credit   int
	Currency string
}

func (a *Account) NewEntry(debit int, credit int, currency string) (*Entry, error) {

	if currency == "" {
		return nil, fmt.Errorf("currency is needed to create an entry")
	}

	e := &Entry{a, debit, credit, currency}
	return e, nil
}
