package preference

type ConfigurationPreference struct {
	Installments        int `json:"installments"`
	DefaultInstallments int `json:"default_installments"`
}
