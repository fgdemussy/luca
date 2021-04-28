package ledger

import "testing"

func TestAccount_GetName(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"should return name", fields{"Expenses"}, "Expenses"},
		{"should return empty string", fields{}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				Name: tt.fields.Name,
			}
			if got := a.GetName(); got != tt.want {
				t.Errorf("Account.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
