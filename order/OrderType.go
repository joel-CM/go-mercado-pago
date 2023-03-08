package order

import "github.com/joel-CM/go_mercadopago/preference"

type Order struct {
	PreferenceId      string            `json:"preference_id"`
	ExternalReference string            `json:"external_reference"`
	AdditionalInfo    string            `json:"additional_info"`
	Items             []preference.Item `json:"items"`
	Payer             PayerOrder        `json:"payer"`
}
