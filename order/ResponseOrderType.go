package order

import "github.com/joel-CM/go_mercadopago/preference"

type ResponseOrder struct {
	AdditionalInfo    string            `json:"additional_info"`
	Cancelled         bool              `json:"cancelled"`
	DateCreated       string            `json:"date_created"`
	ExternalReference string            `json:"external_reference"`
	ID                int               `json:"id"`
	Items             []preference.Item `json:"items"`
	LastUpdated       string            `json:"last_updated"`
	OrderStatus       string            `json:"order_status"`
	Payer             PayerOrder        `json:"payer"`
	PreferenceId      string            `json:"preference_id"`
	PaidAmount        float32           `json:"paid_amount"`
	RefundedAmount    int               `json:"refunded_amount"`
	Status            string            `json:"status"`
	TotalAmount       float32           `json:"total_amount"`
}
