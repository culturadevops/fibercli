package driver

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDriver struct {
	d      *Driversql
	Client *gorm.DB
}

func (t *MysqlDriver) GetClient() interface{} {
	return t.Client
}
func (t *MysqlDriver) Configure(ConfigPath string, ConfigName string) {
	x := &Driversql{}
	x.Configure(ConfigPath, ConfigName)
	t.d = x
}

func (c *MysqlDriver) NewDriver() error {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.d.User, c.d.Password, c.d.Host, c.d.Port, c.d.Database, c.d.Charset)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
		return err
	}
	c.Client = db
	return nil
}
