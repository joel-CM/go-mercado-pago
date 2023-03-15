package payment

import "log"

type ResponsePayment struct {
	ID                  int     `json:"id"`
	TransactionAmount   float32 `json:"transaction_amount"`
	TotalPaidAmount     float32 `json:"total_paid_amount"`
	CurrencyId          string  `json:"currency_id"`
	Status              string  `json:"status"`
	StatusDetail        string  `json:"status_detail"`
	OperationType       string  `json:"operation_type"`
	DateApproved        string  `json:"date_approved"`
	DateCreated         string  `json:"date_created"`
	LastModified        string  `json:"last_modified"`
	AmountRefunded      float32 `json:"amount_refunded"`
	CollectorId         string  `json:"collector_id"`
	Description         string  `json:"description"`
	LiveMode            bool    `json:"live_mode"`
	AuthorizationCode   string  `json:"authorization_code"`
	TransactionDetails  string  `json:"transaction_details"`
	BinaryMode          bool    `json:"binary_mode"`
	StatementDescriptor string  `json:"statement_descriptor"`
	Order               struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}
}
