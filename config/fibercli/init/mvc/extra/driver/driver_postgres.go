package driver

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DriverPostgres struct {
	d      *Driversql
	Client *gorm.DB
}

func (t *DriverPostgres) GetClient() interface{} {
	return t.Client
}
func (t *DriverPostgres) Configure(ConfigPath string, ConfigName string) {
	x := &Driversql{}
	x.Configure(ConfigPath, ConfigName)
	t.d = x
}

func (c *DriverPostgres) NewDriver() error {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s", c.d.User, c.d.Password, c.d.Host, c.d.Port, c.d.Database, c.d.TimeZone)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
		os.Exit(-1)
		return err
	}
	c.Client = db
	return nil
}
