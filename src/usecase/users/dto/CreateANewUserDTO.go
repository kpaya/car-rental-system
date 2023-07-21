package dto

import "github.com/kpaya/car-rental-system/src/entity"

type InputCreateANewUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type OutputCreateANewUserDTO struct {
	Id     string
	Name   string
	Email  string
	Status entity.AccountStatus
	Phone  string
}
