package ledger

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestNewEntry(t *testing.T) {
	t.Run("should return a new Entry with provided Account", func(t *testing.T) {
		c := qt.New(t)
		a := &Account{}
		got, err := a.NewEntry(0, 0, "USD")
		c.Assert(got, qt.DeepEquals, &Entry{a, 0, 0, "USD"})
		c.Assert(err, qt.IsNil)
	})
	t.Run("should provide a currency", func(t *testing.T) {
		c := qt.New(t)
		a := &Account{}
		got, err := a.NewEntry(0, 0, "")
		c.Assert(got, qt.IsNil)
		c.Assert(err.Error(), qt.Contains, "currency is needed")
	})
}
