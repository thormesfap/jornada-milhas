package models

import (
	"gorm.io/gorm"
	"gopkg.in/validator.v2"
)

type Depoimento struct {
	
	gorm.Model
	Foto       string `json:"foto" validate:"nonzero"`
	Depoimento string `json:"depoimento" validate:"min=20"`
	Autor      string `json:"autor" validate:"nonzero"`
}

func ValidateDepoimento(depoimento *Depoimento)error{
	if err := validator.Validate(depoimento); err != nil {
		return err
	}
	return nil
}
