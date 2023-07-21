package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type BarcodeReader struct {
	ID        string    `validate:"uuid" json:"id"`
	CreatedAt time.Time `validate:"datetime" json:"created_at"`
	Active    bool
}

func NewBarcodeReader(id string) *BarcodeReader {
	barcodeReader := new(BarcodeReader)
	if id == "" {
		barcodeReader.ID = uuid.NewString()
	}
	barcodeReader.Active = true
	barcodeReader.CreatedAt = time.Now().Local()
	return barcodeReader
}

func (b *BarcodeReader) Validate() error {
	err := validator.New().Struct(b)
	if err != nil {
		return err
	}
	return nil
}
