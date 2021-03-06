package ledger

type SubscriptionType string

const (
	ACCOUNT_BALANCE_LIMIT                   SubscriptionType = "ACCOUNT_BALANCE_LIMIT"
	OFFCHAIN_WITHDRAWAL                                      = "OFFCHAIN_WITHDRAWAL"
	TRANSACTION_HISTORY_REPORT                               = "TRANSACTION_HISTORY_REPORT"
	ACCOUNT_INCOMING_BLOCKCHAIN_TRANSACTION                  = "ACCOUNT_INCOMING_BLOCKCHAIN_TRANSACTION"
	COMPLETE_BLOCKCHAIN_TRANSACTION                          = "COMPLETE_BLOCKCHAIN_TRANSACTION"
)
