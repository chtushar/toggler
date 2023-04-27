package configs

type (
	DB struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}

	Config struct {
		Port       int
		Production bool
		DB         *DB
	}
)
