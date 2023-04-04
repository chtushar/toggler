package cmd

import (
	"flag"
	"fmt"
	"os"
)

func Execute() {
	initConfigPtr := flag.Bool("init-config", false, "Initialize the configuration file")

	flag.Parse()

	if *initConfigPtr {
		fmt.Println("Initializing configuration file")
		os.Exit(0)
	}

	fmt.Println("Hello, world!")
}
