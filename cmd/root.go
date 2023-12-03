package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/chtushar/toggler/adapters/node"
	"github.com/chtushar/toggler/api"
	"github.com/chtushar/toggler/api/app"
	"github.com/chtushar/toggler/configs"
	"github.com/chtushar/toggler/db"
	"github.com/chtushar/toggler/db/queries"
	"github.com/jackc/pgx/v4/stdlib"
)

func Execute(ctx context.Context) {
	// Pointer to the init-config flag
	initConfigPtr := flag.Bool("init-config", false, "Initialize the configuration file")
	runUPMigrationPtr := flag.Bool("up-migration", false, "Run the up migration")
	runDownMigrationPtr := flag.Bool("down-migration", false, "Run the down migration")

	flag.Parse()

	// If the init-config flag is set, create a config file and exit
	if *initConfigPtr {
		err := configs.NewConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	getConfigFromFile()
	
	// Initialize the database
	dbConn = initDB()

	app := &app.App{
		Port:   cfg.Port,
		DbConn: dbConn,
		Jwt:    cfg.JWTSecret,
		Q:      queries.New(dbConn),
		Log:    log.New(os.Stdout, "toggler: ", log.LstdFlags),
	};

	defer dbConn.Close()

	// If the up-migration flag is set, run the up migration and exit
	if *runUPMigrationPtr {
		pgxPoolConfig, err := getPGXConfig()
		if err != nil {
			app.Log.Fatal("Failed to get database config", err)
			os.Exit(1)
		}

		db.RunUpMigrations(stdlib.OpenDB(*pgxPoolConfig.ConnConfig), app.Log)
		os.Exit(0)
	}

	// If the down-migration flag is set, run the down migration and exit
	if *runDownMigrationPtr {
		pgxPoolConfig, err := getPGXConfig()
		if err != nil {
			app.Log.Fatal("Failed to get database config", err)
			os.Exit(1)
		}
		db.RunDownMigration(stdlib.OpenDB(*pgxPoolConfig.ConnConfig), app.Log)
		os.Exit(0)
	}

	n := node.Node{}
	n.Init(ctx)

	app.Node = &n

	// Initialize the HTTP server
	api.InitHTTPServer(app, ctx)
}
