package dto

import "github.com/kpaya/car-rental-system/src/entity"

type InputCreateANewUserDTO struct {
	Name     string
	Email    string
	Password string
	Status   entity.AccountStatus
	Phone    string
}

type OutputCreateANewUserDTO struct {
	Id       string
	Name     string
	Email    string
	Password string
	Status   entity.AccountStatus
	Phone    string
}
