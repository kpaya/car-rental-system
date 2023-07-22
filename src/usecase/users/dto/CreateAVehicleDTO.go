package dto

import "github.com/kpaya/car-rental-system/src/entity"

type InputCreateAVehicleDTO struct {
	SegmentCar        string `json:"segment_car"`
	LicenseNumber     string `json:"license_number"`
	StockNumber       string `json:"stock_number"`
	PassengerCapacity int32  `json:"passenger_capacity"`
	Barcode           string `json:"barcode"`
	HasSunroof        bool   `json:"has_sunroof"`
	Status            string `json:"status"`
	Model             string `json:"model"`
	ManufacturingYear int64  `json:"manufacturing_year"`
	Mileage           int64  `json:"mileage"`
}

type OutputCreateAVehicleDTO struct {
	ID                string               `json:"id"`
	SegmentCar        entity.SegmentCar    `json:"segment_car"`
	LicenseNumber     string               `json:"license_number"`
	StockNumber       string               `json:"stock_number"`
	PassengerCapacity int32                `json:"passenger_capacity"`
	Barcode           string               `json:"barcode"`
	HasSunroof        bool                 `json:"has_sunroof"`
	Status            entity.VehicleStatus `json:"status"`
	Model             string               `json:"model"`
	ManufacturingYear int64                `json:"manufacturing_year"`
	Mileage           int64                `json:"mileage"`
}
