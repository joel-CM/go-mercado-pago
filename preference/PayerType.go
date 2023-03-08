package preference

type PayerPreference struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Phone   struct {
		AreaCode string `json:"area_code"`
		Number   string `json:"number"`
	}
	Identification struct {
		Type   string `json:"type"`   // tipo de identification
		Number string `json:"number"` //
	}
	Address struct {
		ZipCode      string `json:"zip_code"`      //codigo postal
		StreetName   string `json:"street_name"`   // calle
		StreetNumber int    `json:"street_number"` // numero
	}
	DateCreated string `json:"date_created"`
}
