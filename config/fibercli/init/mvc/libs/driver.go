package libs

import (
	"fmt"
	"log"
	"os"

	config "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Driversql struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	TimeZone     string
	print_log    bool
	Client       *gorm.DB
}

func (t *Driversql) Configure(ConfigPath string, ConfigName string) {
	config.AddConfigPath(ConfigPath)
	config.SetConfigName(ConfigName)
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if config.GetString("default.host") == "" {
		fmt.Println("falta host en el archivo " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("default.database") == "" {
		fmt.Println("falta database en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("default.user") == "" {
		fmt.Println("falta user en el archivo de config " + ConfigName)
		os.Exit(1)
	}

	t.Host = config.GetString("default.host")
	t.Port = config.GetString("default.port")
	t.Database = config.GetString("default.database")
	t.User = config.GetString("default.user")
	t.Password = config.GetString("default.password")
	t.Charset = config.GetString("default.charset")
	t.MaxIdleConns = config.GetInt("default.MaxIdleConns")
	t.MaxOpenConns = config.GetInt("default.MaxOpenConns")
	t.print_log = config.GetBool("default.sql_log")

}

type MysqlDriver struct {
	d      *Driversql
	Client *gorm.DB
}

func (t *MysqlDriver) GetClient() *gorm.DB {
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
