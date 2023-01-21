package usecase

import (
	"errors"

	repository_interfaces "github.com/kpaya/car-rental-system/src/repository/interfaces"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
)

type FindAUserByIdUseCase struct {
	Repository repository_interfaces.UserRepositoryInterface
}

func NewFindAUserByIdUseCase(repository repository_interfaces.UserRepositoryInterface) FindAUserByIdUseCase {
	return FindAUserByIdUseCase{Repository: repository}
}

func (u *FindAUserByIdUseCase) Execute(input dto.InputFindAUserByIdDTO) (dto.OutputFindAUserByIdDTO, error) {
	foundUser, _ := u.Repository.FindById(input.Id)
	if foundUser.ID == "" {
		return dto.OutputFindAUserByIdDTO{}, errors.New("this user doesn't exists")
	}
	output := dto.OutputFindAUserByIdDTO{
		Id:     foundUser.ID,
		Name:   foundUser.Name,
		Email:  foundUser.Email,
		Status: foundUser.Status,
		Phone:  foundUser.Phone,
	}

	return output, nil
}
