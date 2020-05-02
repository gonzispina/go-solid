package wrong

type Type int
const (
	National Type = iota
	International
)

// Creates a new transaction
func New(t Type, amount int) Transaction {
	return Transaction{
		t:      t,
		amount: amount,
	}
}

// Transaction implementation
type Transaction struct {
	t Type
	amount int
}

// Amount returns the amount of the transaction
func (t *Transaction) Amount() int {
	return t.amount
}

// Cancel the transaction if possible
func (t *Transaction) Cancel() bool {
	if t.t == International {
		return false
	}

	// Do something
	return true
}

