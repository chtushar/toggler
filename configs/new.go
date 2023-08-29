package configs

import (
	"embed"
	"fmt"
	"os"
)

var ConfigExample embed.FS

func NewConfig() error {
	content, _ := ConfigExample.ReadFile("config-example.yaml")
	
	f, err := os.Create("config.yaml")

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(content)

	if err != nil {
		return err
	}

	f.Sync()
	fmt.Println("Initiated new config at: config.yaml")

	return nil
}
