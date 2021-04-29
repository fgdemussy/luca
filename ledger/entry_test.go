package ledger

import (
	"testing"
)

func TestNewEntry(t *testing.T) {
	a := &Account{}
	got, err := a.NewEntry(0, 0, "USD")
	if err != nil {
		t.Errorf("Should've created a NewEntry: %v, but got error: %v", got, err)
	}
	if got == nil {
		t.Error("Should've created a NewEntry, but got nil")
	}

	got, err = a.NewEntry(0, 0, "")
	if err == nil {
		t.Errorf("Should've failed creating NewEntry, but got: %v", got)
	}
}
