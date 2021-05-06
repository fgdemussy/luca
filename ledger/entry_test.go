package ledger

import (
	"reflect"
	"testing"

	"github.com/Rhymond/go-money"
)

func TestAccount_NewEntry(t *testing.T) {
	type fields struct {
		Name string
	}
	type args struct {
		operation string
		amount    int64
		code      string
	}
	type test struct {
		name string
		fields
		args
		want    *Entry
		wantErr bool
	}
	tests := []test{
		{
			"Fails given uknown operation",
			fields{"Account A"},
			args{"boom", 0, "USD"},
			nil,
			true,
		},
		{
			"Fails given amount < 0",
			fields{"Account A"},
			args{"debit", -10, "USD"},
			nil,
			true,
		},
		{
			"Fails given unknown currency code",
			fields{"Account A"},
			args{"debit", 0, "ZAP"},
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{tt.fields.Name}
			got, err := a.NewEntry(tt.args.operation, tt.args.amount, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("Account.NewEntry error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Account.NewEntry got = %v, want = %v", got, tt.want)
			}
		})
	}

	t.Run("Returns a new Entry", func(t *testing.T) {
		var amount int64
		amount = 10
		a := &Account{"A"}
		code := "USD"
		m := money.New(amount, code)
		o := &Operation{debit}
		want := &Entry{Account: a, Operation: o, Money: m}
		got, err := a.NewEntry(debit, amount, code)
		if err != nil {
			t.Errorf("Account.NewEntry fails with error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Account.NewEntry got = %v, want = %v", got, want)
		}
	})
}
