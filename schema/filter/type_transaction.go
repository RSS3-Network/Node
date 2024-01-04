package filter

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=TransactionType --transform=snake --trimprefix=TypeTransaction --output type_transaction_string.go --json --sql
type TransactionType uint64

//goland:noinspection GoMixedReceiverTypes
func (i TransactionType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i TransactionType) Tag() Tag {
	return TagTransaction
}

const (
	TypeTransactionTransfer TransactionType = iota + 1
	TypeTransactionApproval
	TypeTransactionBurn
	TypeTransactionMint
	TypeTransactionBridge
)
