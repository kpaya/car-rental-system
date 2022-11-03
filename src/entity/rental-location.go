package entity

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kpaya/car-rental-system/src/entity/value_object"
)

type RentalLocation struct {
	ID      string `validate:"uuid4"`
	Name    string `validate:"required"`
	Address value_object.Location
}

func NewRentalLocation(id string, name string, address value_object.Location) *RentalLocation {
	var rentalLocation = new(RentalLocation)
	if id == "" {
		rentalLocation.ID = uuid.NewString()
	}
	rentalLocation.Name = name
	rentalLocation.Address = address
	return rentalLocation
}

func (l *RentalLocation) Validate() {
	err := validator.New().Struct(l)
	if err != nil {
		log.Panic(err)
	}
}
