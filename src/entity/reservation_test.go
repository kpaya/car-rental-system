package entity_test

import (
	"testing"
	"time"

	"github.com/kpaya/car-rental-system/src/entity"
	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	assert := assert.New(t)

	reservation, err := entity.NewReservation(entity.NewVehicle("", entity.Car, "23132SS", "1B", 4, "", true, "Pajeiro Dakar", 2022, 200), time.Now().Add(time.Duration(200)), "Centro", "Caixa")
	assert.Nil(err)
	assert.NotNil(reservation)
}
