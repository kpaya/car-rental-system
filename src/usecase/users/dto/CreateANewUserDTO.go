package dto

import (
	"github.com/kpaya/car-rental-system/src/entity"
	"github.com/kpaya/car-rental-system/src/entity/value_object"
)

type InputCreateANewUserDTO struct {
	Name     string               `json:"name"`
	Email    string               `json:"email"`
	Password string               `json:"password"`
	Phone    string               `json:"phone"`
	Address  value_object.Address `json:"address"`
}

type OutputCreateANewUserDTO struct {
	Id      string
	Name    string
	Email   string
	Status  entity.AccountStatus
	Phone   string
	Address value_object.Address `json:"address"`
}
