package preference

type Preference struct {
	Items           []Item          `json:"items"`
	BackUrls        BackUrls        `json:"back_urls"`
	NotificationUrl string          `json:"notification_url"`
	Payer           PayerPreference `json:"payer"`
	PaymentMethods  struct {
		Installments        int `json:"installments"`
		DefaultInstallments int `json:"default_installments"`
	} `json:"payment_methods"`
}
