package preference

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	mercadopago "github.com/joel-CM/go_mercadopago"
)

func (p Preference) Create(configurationPreference ConfigurationPreference) (ResponsePreference, error) {
	setConfigurationPreference(&p, configurationPreference)

	var client *http.Client = &http.Client{}
	var respPref ResponsePreference = ResponsePreference{}

	jsonPreference, jsonPreferenceErr := json.Marshal(p)
	if jsonPreferenceErr != nil {
		return respPref, jsonPreferenceErr
	}

	req, reqErr := http.NewRequest(
		"POST",
		mercadopago.ApiPreferenceMP,
		bytes.NewBuffer(jsonPreference),
	)
	if reqErr != nil {
		return respPref, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", mercadopago.AccessToken))

	resp, respErr := client.Do(req)
	if respErr != nil {
		return respPref, respErr
	}

	defer resp.Body.Close()

	decodeErr := json.NewDecoder(resp.Body).Decode(&respPref)
	if decodeErr != nil {
		return respPref, decodeErr
	}

	return respPref, nil
}

func setConfigurationPreference(p *Preference, configurationPreference ConfigurationPreference) {
	if configurationPreference.Installments > 0 {
		p.PaymentMethods.Installments = configurationPreference.Installments
	} else {
		p.PaymentMethods.Installments = 1
	}
}
