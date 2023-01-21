package repository

import (
	"database/sql"
	"errors"

	"github.com/kpaya/car-rental-system/src/entity"
)

type VehicleRepository struct {
	DB *sql.DB
}

func NewVehicleRepository(db *sql.DB) *VehicleRepository {
	return &VehicleRepository{
		DB: db,
	}
}

func (v *VehicleRepository) Create(vehicle *entity.Vehicle) error {
	stmt, err := v.DB.Prepare("INSERT INTO vehicle (id, segment_car, license_number, stock_number, passenger_capacity, barcode, has_sunroof, status, model, manufacturing_year, milage) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	var convertSunroof int8
	if vehicle.HasSunroof {
		convertSunroof = 1
	} else {
		convertSunroof = 0
	}
	_, err = stmt.Exec(vehicle.ID, vehicle.SegmentCar, vehicle.LicenseNumber, vehicle.StockNumber, vehicle.PassengerCapacity, vehicle.Barcode, convertSunroof, vehicle.Status, vehicle.Model, vehicle.ManufacturingYear, vehicle.Mileage)
	if err != nil {
		return err
	}
	return nil
}

func (v *VehicleRepository) Update(vehicle *entity.Vehicle) error {
	if v.FindById(vehicle.ID).ID == "" {
		return errors.New("we couldn't be able to find the record")
	}
	stmt, err := v.DB.Prepare("UPDATE vehicle SET segment_car = $1, license_number = $2, stock_number = $3, passenger_capacity = $4, barcode = $5, has_sunroof = $6, status = $7, model = $8, manufacturing_year = $9, milage = $10")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(vehicle.SegmentCar, vehicle.LicenseNumber, vehicle.StockNumber, vehicle.PassengerCapacity, vehicle.Barcode, vehicle.HasSunroof, vehicle.Status, vehicle.Model, vehicle.ManufacturingYear, vehicle.Mileage)
	if err != nil {
		return err
	}
	return nil
}

func (v *VehicleRepository) FindById(id string) *entity.Vehicle {
	var vehicle = new(entity.Vehicle)
	row := v.DB.QueryRow("SELECT * FROM vehicle WHERE id = $1", id)

	err := row.Scan(&vehicle.ID, &vehicle.SegmentCar, &vehicle.LicenseNumber, &vehicle.StockNumber, &vehicle.PassengerCapacity, &vehicle.Barcode, &vehicle.HasSunroof, &vehicle.Status, &vehicle.Model, &vehicle.ManufacturingYear, &vehicle.Mileage)
	if err != nil {
		return &entity.Vehicle{}
	}
	return vehicle
}

func (v *VehicleRepository) List() []*entity.Vehicle {
	var listVehicles []*entity.Vehicle
	rows, err := v.DB.Query("SELECT * FROM vehicle")
	if err != nil {
		return listVehicles
	}
	defer rows.Close()
	for rows.Next() {
		var vehicle *entity.Vehicle
		err := rows.Scan(&vehicle.ID, &vehicle.SegmentCar, &vehicle.LicenseNumber, &vehicle.StockNumber, &vehicle.PassengerCapacity, &vehicle.Barcode, &vehicle.HasSunroof, &vehicle.Status, &vehicle.Model, &vehicle.ManufacturingYear, &vehicle.Mileage)
		if err != nil {
			return []*entity.Vehicle{}
		}
		listVehicles = append(listVehicles, vehicle)
	}
	return listVehicles
}
