package usecase

import (
	"github.com/kpaya/car-rental-system/src/entity"
	repository_interfaces "github.com/kpaya/car-rental-system/src/repository/interfaces"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
)

type CreateAVehicleUseCase struct {
	Repository repository_interfaces.VehicleRepositoryInterface
}

func NewCreateVehicleUseCase(repository repository_interfaces.VehicleRepositoryInterface) *CreateAVehicleUseCase {
	return &CreateAVehicleUseCase{
		Repository: repository,
	}
}

func (v *CreateAVehicleUseCase) Execute(input dto.InputCreateAVehicleDTO) dto.OutputCreateAVehicleDTO {
	vehicle := entity.NewVehicle("", entity.SegmentCar(input.SegmentCar), input.LicenseNumber, input.StockNumber, input.PassengerCapacity, input.Barcode, input.HasSunroof, input.Model, input.ManufacturingYear, input.Mileage)

	if err := v.Repository.Create(vehicle); err != nil {
		return dto.OutputCreateAVehicleDTO{}
	}

	output := dto.OutputCreateAVehicleDTO{
		ID:                vehicle.ID,
		SegmentCar:        vehicle.SegmentCar,
		LicenseNumber:     vehicle.LicenseNumber,
		StockNumber:       vehicle.StockNumber,
		PassengerCapacity: vehicle.PassengerCapacity,
		Barcode:           vehicle.Barcode,
		HasSunroof:        vehicle.HasSunroof,
		Status:            vehicle.Status,
		Model:             vehicle.Model,
		ManufacturingYear: vehicle.ManufacturingYear,
		Mileage:           vehicle.Mileage,
	}

	return output
}
