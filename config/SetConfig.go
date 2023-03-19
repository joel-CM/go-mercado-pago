package config

var Config Configuration

func Set(c Configuration) {
	Config.AccessToken = c.AccessToken
	Config.ApiPreferenceMP = c.ApiPreferenceMP
	Config.ApiOrderMP = c.ApiOrderMP
	Config.ApiPaymentMP = c.ApiPaymentMP
}
