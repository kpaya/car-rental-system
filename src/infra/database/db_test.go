package database_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestCreatingConnectionWithDb(t *testing.T) {
	_ = godotenv.Load(".env")

	assert := assert.New(t)

	assert.Equal(os.Getenv("DB_DRIVER"), "Tubs")
}
