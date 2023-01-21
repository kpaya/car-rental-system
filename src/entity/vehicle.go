package entity

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type VehicleStatus string

const (
	Avaliable     VehicleStatus = "Avaliable"
	Reserved      VehicleStatus = "Reserved"
	Loaned        VehicleStatus = "Loaned"
	Lost          VehicleStatus = "Lost"
	BeingServiced VehicleStatus = "BeingServiced"
	Other         VehicleStatus = "Other"
)

type SegmentCar string

const (
	Car        SegmentCar = "Car"
	Truck      SegmentCar = "Truck"
	SUV        SegmentCar = "SUV"
	Van        SegmentCar = "Van"
	Motorcycle SegmentCar = "Motorcycle"
)

type Vehicle struct {
	ID                string `validate:"uuid4"`
	SegmentCar        SegmentCar
	LicenseNumber     string
	StockNumber       string
	PassengerCapacity int32
	Barcode           string
	HasSunroof        bool
	Status            VehicleStatus
	Model             string
	ManufacturingYear int64
	Mileage           int64
}

func NewVehicle(id string, segmentCar SegmentCar, licenseNumber string, stockNumber string, passengerCapacity int32, barcode string, hasSunroof bool, model string, manufacturingYear int64, mileage int64) *Vehicle {
	var vehicle = new(Vehicle)
	if id == "" {
		vehicle.ID = uuid.NewString()
	}
	vehicle.SegmentCar = segmentCar
	vehicle.LicenseNumber = licenseNumber
	vehicle.StockNumber = stockNumber
	vehicle.PassengerCapacity = passengerCapacity
	vehicle.Barcode = barcode
	vehicle.HasSunroof = hasSunroof
	vehicle.Status = Avaliable
	vehicle.Model = model
	vehicle.ManufacturingYear = manufacturingYear
	vehicle.Mileage = mileage

	return vehicle
}

func (l *Vehicle) Validate() {
	err := validator.New().Struct(l)
	if err != nil {
		log.Panic(err)
	}
}
