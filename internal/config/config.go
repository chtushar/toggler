package config

import (
	"log"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  *App
	once sync.Once
)

func Get() *App {
	readConfig()

	once.Do(func() {
		cfg = &App{
			Port:       8080,
			Production: false,
			DB: &DB{
				Host:     "localhost",
				Port:     5432,
				User:     "postgres",
				Password: "postgres",
				Name:     "postgres",
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
