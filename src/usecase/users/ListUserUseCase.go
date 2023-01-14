package usecase

import (
	"github.com/kpaya/car-rental-system/src/repository"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
)

type ListUserUseCase struct {
	Repository repository.UserRepository
}

func NewListUserUseCase(repository *repository.UserRepository) ListUserUseCase {
	return ListUserUseCase{
		Repository: *repository,
	}
}

func (u *ListUserUseCase) Execute() (dto.OutputListUserDTO, error) {
	users, err := u.Repository.List()
	if err != nil {
		return dto.OutputListUserDTO{}, err
	}
	return dto.OutputListUserDTO{
		List: users,
	}, nil
}
