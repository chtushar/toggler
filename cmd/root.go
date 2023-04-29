package cmd

import (
	"flag"
	"fmt"
	"os"
)

func Execute() {
	// Pointer to the init-config flag
	initConfigPtr := flag.Bool("init-config", false, "Initialize the configuration file")

	flag.Parse()

	// If the init-config flag is set, create a config file and exit
	if *initConfigPtr {
		err := newConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}
