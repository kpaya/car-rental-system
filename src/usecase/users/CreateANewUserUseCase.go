package usecase

import (
	"github.com/kpaya/car-rental-system/src/entity"
	repository_interfaces "github.com/kpaya/car-rental-system/src/repository/interfaces"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
)

type CreateANewUserUseCase struct {
	Repository repository_interfaces.RepositoryInterface
}

func NewCreateANewUserUseCase(repository repository_interfaces.RepositoryInterface) CreateANewUserUseCase {
	return CreateANewUserUseCase{Repository: repository}
}

func (u *CreateANewUserUseCase) Execute(input dto.InputCreateANewUserDTO) (dto.OutputCreateANewUserDTO, error) {
	user := entity.NewUser("", input.Name, input.Password, entity.Active, input.Email, input.Phone)
	err := u.Repository.Create(user)
	if err != nil {
		return dto.OutputCreateANewUserDTO{}, err
	}
	output := dto.OutputCreateANewUserDTO{
		Id:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}

	return output, nil
}
