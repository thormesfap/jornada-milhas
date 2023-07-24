package models

import "gorm.io/gorm"

type Depoimento struct {
	
	gorm.Model
	Foto       string `json:"foto"`
	Depoimento string `json:"depoimento"`
	Autor      string `json:"autor"`
}
