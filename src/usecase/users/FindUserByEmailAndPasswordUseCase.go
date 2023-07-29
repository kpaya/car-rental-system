package usecase

import (
	"errors"

	repository_interfaces "github.com/kpaya/car-rental-system/src/repository/interfaces"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
)

type FindUserByEmailAndPasswordUseCase struct {
	Repository repository_interfaces.UserRepositoryInterface
}

func NewFindUserByEmailAndPasswordUseCase(repository repository_interfaces.UserRepositoryInterface) *FindUserByEmailAndPasswordUseCase {
	return &FindUserByEmailAndPasswordUseCase{
		Repository: repository,
	}
}

func (u *FindUserByEmailAndPasswordUseCase) Execute(input dto.InputFindUserByEmailAndPasswordDTO) (dto.OutputFindUserByEmailAndPasswordDTO, error) {
	user, err := u.Repository.FindUserByEmailAndPassword(input.Email, input.Password)
	if err != nil {
		return dto.OutputFindUserByEmailAndPasswordDTO{}, errors.New("we cannot find a user according with provided data")
	}
	return dto.OutputFindUserByEmailAndPasswordDTO{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Type:   user.Type,
		Status: string(user.Status),
	}, nil
}
