package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/chtushar/toggler/db"
	"github.com/jackc/pgx/v4/stdlib"
)

func Execute() {
	// Pointer to the init-config flag
	initConfigPtr := flag.Bool("init-config", false, "Initialize the configuration file")
	runUPMigrationPtr := flag.Bool("up-migration", false, "Run the up migration")
	runDownMigrationPtr := flag.Bool("down-migration", false, "Run the down migration")

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

	// Initialize the database
	dbConn = initDB()

	app := &App{
		port:   cfg.Port,
		dbConn: dbConn,
		log:    log.New(os.Stdout, "toggler: ", log.LstdFlags),
	}

	// If the up-migration flag is set, run the up migration and exit
	if *runUPMigrationPtr {
		pgxConfig, err := getPGXConfig()
		if err != nil {
			app.log.Fatal("Failed to get database config", err)
			os.Exit(1)
		}

		db.RunUpMigrations(stdlib.OpenDB(*pgxConfig), app.log)
		os.Exit(0)
	}

	// If the down-migration flag is set, run the down migration and exit
	if *runDownMigrationPtr {
		pgxConfig, err := getPGXConfig()
		if err != nil {
			app.log.Fatal("Failed to get database config", err)
			os.Exit(1)
		}
		db.RunDownMigration(stdlib.OpenDB(*pgxConfig), app.log)
		os.Exit(0)
	}

	// Initialize the HTTP server
	initHTTPServer(app)
}
