package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Destino struct {
	gorm.Model
	Foto  string
	Nome  string  `json:"nome" validate:"nonzero"`
	Preco float64 `json:"preco" validate:"nonzero"`
}

func ValidateDestino(destino *Destino) error {
	if err := validator.Validate(destino); err != nil {
		return err
	}
	return nil
}
