package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	c "github.com/joel-CM/go_mercadopago/config"
)

func (order *Order) Create() (ResponseOrder, error) {
	var responseOrder ResponseOrder
	var client *http.Client = &http.Client{}

	jsonOrder, _ := json.Marshal(order)

	req, reqErr := http.NewRequest("POST", c.Config.ApiOrderMP, bytes.NewBuffer(jsonOrder))
	if reqErr != nil {
		return responseOrder, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Config.AccessToken))

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

func GetById(id int) (ResponseOrder, error) {
	var client *http.Client = &http.Client{}
	var order ResponseOrder

	req, reqErr := http.NewRequest("GET", fmt.Sprintf("%v/%v", c.Config.ApiOrderMP, id), nil)
	if reqErr != nil {
		return order, reqErr
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Config.AccessToken))

	respOrder, respOrderErr := client.Do(req)
	if respOrderErr != nil {
		return order, respOrderErr
	}

	defer respOrder.Body.Close()

	decErr := json.NewDecoder(respOrder.Body).Decode(&order)
	if decErr != nil {
		return order, decErr
	}

	return order, nil
}
