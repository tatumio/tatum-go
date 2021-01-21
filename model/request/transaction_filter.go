package request

import "github.com/tatumio/tatum-go/model/response/ledger"

type TransactionFilter struct {
	Id              string
	From            uint32
	To              uint32
	Account         string
	CounterAccount  string
	Currency        string
	PaymentId       string
	TransactionCode string
	SenderNote      string
	RecipientNote   string
	OpType          ledger.OperationType
	TransactionType ledger.TransactionType
}
