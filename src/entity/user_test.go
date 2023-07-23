package entity_test

import (
	"testing"

	"github.com/kpaya/car-rental-system/src/entity"
	"github.com/kpaya/car-rental-system/src/entity/value_object"
	"github.com/stretchr/testify/assert"
)

func TestCreateANewUser(t *testing.T) {
	assert := assert.New(t)

	user, _ := entity.NewUser("", "Donderio De Souza", "Olá", entity.Active, "test@gmail.com", "11988231123", value_object.Address{
		ID:            "4234324",
		StreetAddress: "Rua dos Bobos",
		City:          "São Paulo",
		State:         "SP",
		Zipcode:       "12345678",
		Country:       "Brasil",
	})

	assert.Equal("Donderio De Souza", user.Name)
}
