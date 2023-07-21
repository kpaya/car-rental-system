package dto

type InputFindUserByEmailAndPasswordDTO struct {
	Email    string
	Password string
}

type OutputFindUserByEmailAndPasswordDTO struct {
	ID     string
	Name   string
	Email  string
	Status string
}
