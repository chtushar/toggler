package configs

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  *Config
	once sync.Once
)

func Get() *Config {
	readConfig()

	once.Do(func() {
		cfg = &Config{
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
		}
	})
	return cfg
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
		log.Fatal("Failed to read config file")
	}
}

func isProduction(env string) bool {
	return env == "prod"
}
