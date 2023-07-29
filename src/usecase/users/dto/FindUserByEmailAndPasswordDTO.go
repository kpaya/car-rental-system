package dto

type InputFindUserByEmailAndPasswordDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OutputFindUserByEmailAndPasswordDTO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
	Type   string `json:"type"`
}
