package right

type Type int

const (
	National Type = iota
	International
)

type Transaction interface {
	Amount() int
}

// Creates a new transaction
func New(t Type, amount int) Transaction {
}

// BaseTransaction is a basic transaction structure
type BaseTransaction struct {
	t Type
	amount int
}

// Amount retrieves the amount of the transaction
func (b *BaseTransaction) Amount() int { return b.amount }

type InternationalTransaction struct {
	*BaseTransaction
}

type NationalTransaction struct {
	*BaseTransaction
}

func (n *NationalTransaction) Cancel() bool {
	// Do something
	return true
}