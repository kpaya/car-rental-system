package repository_interfaces

import "github.com/kpaya/car-rental-system/src/entity"

type VehicleRepositoryInterface interface {
	Create(vehicle *entity.Vehicle) error
	Update(vehicle *entity.Vehicle) error
	FindById(id string) *entity.Vehicle
	List() []*entity.Vehicle
}
