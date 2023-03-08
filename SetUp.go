package mercadopago

var AccessToken string
var ApiPreferenceMP string
var ApiOrderMP string

func SetUp(access_token string, api_preference_mp string, api_order_mp string) {
	AccessToken = access_token
	ApiPreferenceMP = api_preference_mp
	ApiOrderMP = api_order_mp
}
