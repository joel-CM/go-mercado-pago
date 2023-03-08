package preference

type ResponsePreference struct {
	ID                  string   `json:"id"`
	CollectorId         int      `json:"collector_id"`
	Items               []Item   `json:"items"`
	BackUrls            BackUrls `json:"back_urls"`
	ClientId            string   `json:"client_id"`
	StatementDescriptor string   `json:"statement_descriptor"`
	DateCreated         string   `json:"date_created"`
	InitPoint           string   `json:"init_point"`
}
