package ledger

type Account struct {
	Name string
}

func (a *Account) GetName() string {
	return a.Name
}
