package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ReservationStatus string

const (
	Pending   ReservationStatus = "Pending"
	Confirmed ReservationStatus = "Confirmed"
	Completed ReservationStatus = "Completed"
	Cancelled ReservationStatus = "Cancelled"
)

type Reservation struct {
	ReservationNumber string
	Vehicle           *Vehicle
	Status            ReservationStatus `validate:"required"`
	PickUpLocation    string
	ReturnLocation    string
	CreatedAt         time.Time
	ReturnDate        time.Time
}

func NewReservation(vehicle *Vehicle, returnDate time.Time, pickUpLocation string, returnLocation string) (*Reservation, error) {
	reservation := &Reservation{
		ReservationNumber: uuid.NewString(),
		Vehicle:           vehicle,
		Status:            Pending,
		CreatedAt:         time.Now(),
		ReturnDate:        returnDate,
		PickUpLocation:    pickUpLocation,
		ReturnLocation:    returnLocation,
	}

	err := reservation.Validate()
	if err != nil {
		return &Reservation{}, err
	}
	return reservation, nil
}

func (r *Reservation) Validate() error {
	err := validator.New().Struct(r)
	if err != nil {
		return err
	}
	return nil
}
