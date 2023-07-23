package value_object

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Address struct {
	ID            string `json:"id"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zipcode       string `json:"zip_code"`
	Country       string `json:"contry"`
}

func NewAddress(id string, streetAddress string, city string, state string, zipcode string, country string) (*Address, error) {
	var address = new(Address)
	if id == "" {
		address.ID = uuid.NewString()
	}
	address.StreetAddress = streetAddress
	address.City = city
	address.State = state
	address.Zipcode = zipcode
	address.Country = country

	if err := address.Validate(); err != nil {
		return nil, err
	}

	return address, nil
}

func (l *Address) Validate() error {
	err := validator.New().Struct(l)
	if err != nil {
		return err
	}
	return nil
}
