package ledger

// Account represents a "registry" where Entries are stored.
type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
