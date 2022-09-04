package models

import (
	"github.com/jinzhu/gorm"
)

type Ejemplo struct {
	gorm.Model
	Name string
}

func (t Ejemplo) Holamundo() string {
	var data = "hola mundo"
	return data
}

func (t Ejemplo) ejemploList() []Ejemplo {
	var data = []Ejemplo{}
	return data
}
