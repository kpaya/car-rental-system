package dto

import "github.com/kpaya/car-rental-system/src/entity"

type InputFindAUserByIdDTO struct {
	Id string
}

type OutputFindAUserByIdDTO struct {
	Id     string
	Name   string
	Email  string
	Status entity.AccountStatus
	Phone  string
}
