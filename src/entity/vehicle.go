package entity

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type VehicleStatus string
type CarType string
type VanType string

const (
	Economy      CarType = "Economy"
	Compact      CarType = "Compact"
	Intermediate CarType = "Intermediate"
	Standard     CarType = "Standard"
	FullSize     CarType = "FullSize"
	Premium      CarType = "Premium"
	Luxury       CarType = "Luxury"
)

const (
	Passenger VanType = "Passenger"
	Cargo     VanType = "Cargo"
)

const (
	Avaliable     VehicleStatus = "Avaliable"
	Reserved      VehicleStatus = "Reserved"
	Loaned        VehicleStatus = "Loaned"
	Lost          VehicleStatus = "Lost"
	BeingServiced VehicleStatus = "BeingServiced"
	Other         VehicleStatus = "Other"
)

type Car struct {
	Type CarType
	Vehicle
}

type Truck struct {
	Type string
	Vehicle
}

type SUV struct {
	Type string
	Vehicle
}

type Van struct {
	Type VanType
	Vehicle
}

type Motorcycle struct {
	Type string
	Vehicle
}

type Vehicle struct {
	ID                string `validate:"uuid4"`
	LicenseNumber     string
	StockNumber       string
	PassengerCapacity int32
	Barcode           string
	HasSunroof        bool
	Status            VehicleStatus
	Model             string
	Make              string
	ManufacturingYear int64
	Mileage           int64
}

func NewVehicle(id string) *Vehicle {
	var vehicle = new(Vehicle)
	if id == "" {
		vehicle.ID = uuid.NewString()
	}
	return vehicle
}

func (l *Vehicle) Validate() {
	err := validator.New().Struct(l)
	if err != nil {
		log.Panic(err)
	}
}
