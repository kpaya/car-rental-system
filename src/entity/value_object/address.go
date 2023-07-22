package value_object

import (
	"log"

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

func NewAddress(id string, streetAddress string, city string, state string, zipcode string, country string) *Address {
	var address = new(Address)
	if id == "" {
		address.ID = uuid.NewString()
	}
	address.StreetAddress = streetAddress
	address.City = city
	address.State = state
	address.Zipcode = zipcode
	address.Country = country

	address.Validate()

	return address
}

func (l *Address) Validate() {
	err := validator.New().Struct(l)
	if err != nil {
		log.Panic(err)
	}
}
