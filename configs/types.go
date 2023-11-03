package configs

type (
	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		ForceTLS bool
	}

	Config struct {
		Port       int
		Production bool
		DB         *DB
		JWTSecret  string
		PrivateKey string
	}
)
