package driver

import (
	"fmt"
	"log"
	"os"

	config "github.com/spf13/viper"
	"gorm.io/gorm"
)

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
