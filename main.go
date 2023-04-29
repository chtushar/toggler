package main

import "github.com/chtushar/toggler/cmd"

func main() {
	// Execute with CLI flags
	// No flags start the server
	// init-config flag creates a config file
	cmd.Execute()
}
