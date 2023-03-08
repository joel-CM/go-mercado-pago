package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	mercadopago "github.com/joel-CM/go_mercadopago"
)

func (order *Order) Create() (ResponseOrder, error) {
	var responseOrder ResponseOrder
	var client *http.Client = &http.Client{}

	jsonOrder, _ := json.Marshal(order)

	req, reqErr := http.NewRequest("POST", mercadopago.ApiOrderMP, bytes.NewBuffer(jsonOrder))
	if reqErr != nil {
		return responseOrder, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", mercadopago.AccessToken))

	respOder, respOrderErr := client.Do(req)
	if respOrderErr != nil {
		return responseOrder, respOrderErr
	}

	defer respOder.Body.Close()

	if decErr := json.NewDecoder(respOder.Body).Decode(&responseOrder); decErr != nil {
		return responseOrder, decErr
	}

	return responseOrder, nil
}
