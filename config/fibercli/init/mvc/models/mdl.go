package models

import (
	"github.com/jinzhu/gorm"
)

type Ejemplo struct {
	gorm.Model
	Name string
}

func (t Ejemplo) List() []Ejemplo {
	var data = []Ejemplo{}
	return data
}
