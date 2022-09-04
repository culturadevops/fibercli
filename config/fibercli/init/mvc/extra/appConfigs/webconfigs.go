package appConfigs

import (
	"fmt"
	"log"
	"os"

	config "github.com/spf13/viper"
)

type Web struct {
	Port string
}

func (t *Web) Configure(ConfigPath string, ConfigName string) {
	config.AddConfigPath(ConfigPath)
	config.SetConfigName(ConfigName)
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if config.GetString("default.Port") == "" {
		fmt.Println("falta puerto en el archivo " + ConfigName)
		os.Exit(1)
	}

	t.Port = config.GetString("default.port")

}
