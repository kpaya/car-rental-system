package usecase

import (
	"errors"

	"github.com/kpaya/car-rental-system/src/entity"
	repository_interfaces "github.com/kpaya/car-rental-system/src/repository/interfaces"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
)

type CreateANewUserUseCase struct {
	Repository repository_interfaces.UserRepositoryInterface
}

func NewCreateANewUserUseCase(repository repository_interfaces.UserRepositoryInterface) CreateANewUserUseCase {
	return CreateANewUserUseCase{Repository: repository}
}

func (u *CreateANewUserUseCase) Execute(input dto.InputCreateANewUserDTO) (dto.OutputCreateANewUserDTO, error) {
	userFound := u.Repository.FindByEmail(input.Email)
	if userFound.ID != "" {
		return dto.OutputCreateANewUserDTO{}, errors.New("this email is already in use")
	}
	user, err := entity.NewUser("", input.Name, input.Password, entity.Active, input.Email, input.Phone)
	if err != nil {
		return dto.OutputCreateANewUserDTO{}, err
	}
	err = u.Repository.Create(user)
	if err != nil {
		return dto.OutputCreateANewUserDTO{}, err
	}
	output := dto.OutputCreateANewUserDTO{
		Id:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
		Phone:  user.Phone,
	}

	return output, nil
}
