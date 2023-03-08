package preference

type Item struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	PictureUrl  string          `json:"picture_url"`
	CategoryId  string          `json:"category_id"`
	Quantity    int             `json:"quantity"`
	CurrencyId  string          `json:"currency_id"`
	UnitPrice   float32         `json:"unit_price"`
	Payer       PayerPreference `json:"payer"`
}
