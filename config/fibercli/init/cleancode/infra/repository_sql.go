package infra

import (
	"errors"

	"gorm.io/gorm"
)

type RepositorySql struct {
	d      *Driver
	Client *gorm.DB
}

const (
	ModelErrRecordNotFound  string = "no existe el registro"
	ModelErrRecordNotCreate string = "Registro ya existe no se puede crear nuevamente"
)

func (t *RepositorySql) GetDriver() interface{} {
	return t.Client
}
func (t *RepositorySql) SetDriver(d Driver) error {
	t.d = &d
	t.Client = d.GetClient().(*gorm.DB)
	return nil
}
func (t *RepositorySql) CreateField(filter interface{}, v interface{}) error {
	err := t.GetField(filter)
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
func (t *RepositorySql) GetField(v interface{}) error {
	result := t.Client.Where(v).First(v)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
func (t *RepositorySql) ListAllFor(filters interface{}, v interface{}) error {
	var err error
	if filters == nil {
		err = t.Client.Find(v).Error
	} else {
		err = t.Client.Where(filters).Find(v).Error
	}
	return err
}
