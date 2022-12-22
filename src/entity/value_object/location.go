package value_object

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type Location struct {
	StreetAddress string
	City          string
	State         string
	Zipcode       string
	Country       string
}

func NewAddress(streetAddress string, city string, state string, zipcode string, country string) *Location {
	var location = new(Location)
	location.StreetAddress = streetAddress
	location.City = city
	location.State = state
	location.Zipcode = zipcode
	location.Country = country

	location.Validate()

	return location
}

func (l *Location) Validate() {
	err := validator.New().Struct(l)
	if err != nil {
		log.Panic(err)
	}
}
