package entity

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kpaya/car-rental-system/src/entity/value_object"
)

type RentalAddress struct {
	ID      string `validate:"uuid4"`
	Name    string `validate:"required"`
	Address value_object.Address
}

func NewRentalAddress(id string, name string, address value_object.Address) *RentalAddress {
	var rentalAddress = new(RentalAddress)
	if id == "" {
		rentalAddress.ID = uuid.NewString()
	}
	rentalAddress.Name = name
	rentalAddress.Address = address
	return rentalAddress
}

func (l *RentalAddress) Validate() {
	err := validator.New().Struct(l)
	if err != nil {
		log.Panic(err)
	}
}
