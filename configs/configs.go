package configs

import (
	"log"
	"os"
	"sync"

	"github.com/chtushar/toggler/utils"
	"github.com/spf13/viper"
)

var (
	Cfg  *Config
	once sync.Once
)

func Get() *Config {
	readConfig()

	once.Do(func() {
		Cfg = &Config{
			Port:       viper.GetInt(keyPort),
			Production: isProduction(viper.GetString(keyEnv)),
			DB: &DB{
				Host:     viper.GetString(keyDBHost),
				Port:     viper.GetInt(keyDBPort),
				User:     viper.GetString(keyDBUser),
				Password: viper.GetString(keyDBPass),
				Name:     viper.GetString(keyDBName),
				ForceTLS: viper.GetBool(keyDBForceTLS),
			},
			JWTSecret: viper.GetString(keyJWTSecret),
		}
	})
	return Cfg
}

func readConfig() {
	// Read config from file
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Failed to find home directory")
	}

	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(homeDir)

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Failed to read config.yaml file")
		y := utils.YesNoPrompt("Do you want to create a config file?", false)

		if y {
			NewConfig()
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

func isProduction(env string) bool {
	return env == "prod"
}
