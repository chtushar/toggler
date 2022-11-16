package configs

const (
	defaultPort      = 9090
	defaultEnv       = "prod"
	defaultJWTSecret = "somesecret"
	defaultDBHost    = "localhost"
	defaultDBPort    = 5432
	defaultDBUser    = "postgres"
	defaultDBName    = "toggler"
)

var defaults = map[string]interface{}{
	keyPort:   defaultPort,
	keyEnv:    defaultEnv,
	keyDBHost: defaultDBHost,
	keyDBPort: defaultDBPort,
	keyDBUser: defaultDBUser,
	keyDBName: defaultDBName,
}
