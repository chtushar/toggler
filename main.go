package main

import (
	"embed"

	"github.com/chtushar/toggler/cmd"
	"github.com/chtushar/toggler/configs"
)

//go:embed *.yaml
var configExample embed.FS


func main() {
	// Embed static files
	configs.ConfigExample = configExample
	
	// Execute with CLI flags
	// No flags start the server
	// init-config flag creates a config file
	cmd.Execute()
}
