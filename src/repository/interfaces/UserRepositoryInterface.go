package repository_interfaces

import "github.com/kpaya/car-rental-system/src/entity"

type UserRepositoryInterface interface {
	Create(user *entity.User) error
	Update(user *entity.User) error
	FindById(id string) (entity.User, error)
	FindByEmail(email string) entity.User
	Delete(id string) error
	List() ([]entity.User, error)
	FindUserByEmailAndPassword(email string, password string) (entity.User, error)
}
