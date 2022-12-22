package repository_interfaces

import "github.com/kpaya/car-rental-system/src/entity"

type RepositoryInterface interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	FindById(id string) (entity.User, error)
	Delete(id string) error
	List() ([]entity.User, error)
}
