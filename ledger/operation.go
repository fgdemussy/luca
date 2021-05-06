package ledger

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
