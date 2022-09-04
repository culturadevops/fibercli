package libs

import (
	"errors"

	"gorm.io/gorm"
)

var Query *Modelsql

type Modelsql struct {
	Client *gorm.DB
}

const (
	ModelErrRecordNotFound  string = "no existe el registro"
	ModelErrRecordNotCreate string = "Registro ya existe no se puede crear nuevamente"
)

func (t *Modelsql) GetDriver() *gorm.DB {
	return t.Client
}
func (t *Modelsql) SetDriver(d *gorm.DB) error {
	t.Client = d
	return nil
}
func (t *Modelsql) Create(filter interface{}, v interface{}) error {
	err := t.GetFirstByField(filter)
	if err == nil {
		return errors.New(ModelErrRecordNotCreate)
	}
	if gorm.ErrRecordNotFound == err {
		if err := t.Client.Create(v).Error; err != nil {
			return err
		}
	}
	return err
}

func (t *Modelsql) GetFirstByField(v interface{}) error {
	result := t.Client.Where(v).First(v)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (t *Modelsql) ListAllFor(filters interface{}, v interface{}) error {
	var err error
	if filters == nil {
		err = t.Client.Find(v).Error
	} else {
		err = t.Client.Where(filters).Find(v).Error
	}
	return err
}
func (t *Modelsql) DelBy(v interface{}) error {
	if err := t.Client.Where(v).Delete(v).Error; err != nil {
		return err
	}
	return nil
}

func (t *Modelsql) DelByID(id uint, v interface{}) error {
	if err := t.Client.Where("id = ?", id).Delete(v).Error; err != nil {
		return err
	}
	return nil
}
func (t *Modelsql) GetAllForID(id uint, v interface{}) error {
	err := t.Client.Where("id = ? ", id).Find(v).Error
	if err != nil {
		return err
	}
	return nil
}
func (t *Modelsql) UpdateField(filter interface{}, column string, value interface{}) error {
	err := t.Client.Model(filter).Where(filter).Update(column, value).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *Modelsql) Update(filter interface{}, v interface{}) error {

	err := t.GetFirstByField(filter)
	if err != nil {
		return err
	}
	if err := t.Client.Model(filter).Where(filter).Save(v).Error; err != nil {
		return err
	}
	return nil

}
