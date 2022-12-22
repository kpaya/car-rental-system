package entity_test

import (
	"testing"

	"github.com/kpaya/car-rental-system/src/entity"
	"github.com/kpaya/car-rental-system/src/entity/value_object"
	"github.com/stretchr/testify/assert"
)

func TestCreateANewUser(t *testing.T) {
	assert := assert.New(t)

	user := entity.NewUser("", "Donderio De Souza", "Olá", entity.Active, "test@gmail.com", "11988231123")
	user.Address = *value_object.NewAddress("Avenida Brasil", "São Paulo", "São Paulo", "00000000", "Brazil")

	assert.Equal("Donderio De Souza", user.Name)
}
