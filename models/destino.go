package models

import (
	"gorm.io/gorm"
	"gopkg.in/validator.v2"
)

type Destino struct {
	
	gorm.Model
	Foto       string `json:"foto" validate:"nonzero"`
	Nome string `json:"nome" validate:"nonzero"`
	Preco      float64 `json:"preco" validate:"nonzero"`
}

func ValidateDestino(destino *Destino)error{
	if err := validator.Validate(destino); err != nil {
		return err
	}
	return nil
}
