package dto

import "github.com/kpaya/car-rental-system/src/entity"

type InputCreateAVehicleDTO struct {
	SegmentCar        string
	LicenseNumber     string
	StockNumber       string
	PassengerCapacity int32
	Barcode           string
	HasSunroof        bool
	Status            string
	Model             string
	ManufacturingYear int64
	Mileage           int64
}

type OutputCreateAVehicleDTO struct {
	ID                string
	SegmentCar        entity.SegmentCar
	LicenseNumber     string
	StockNumber       string
	PassengerCapacity int32
	Barcode           string
	HasSunroof        bool
	Status            entity.VehicleStatus
	Model             string
	ManufacturingYear int64
	Mileage           int64
}
