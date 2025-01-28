package configs

type (
	Config struct {
		Service  Service
		Database Database
	}

	Service struct {
		Port      string
		SecretKey string
	}

	Database struct {
		DataSourceName string
	}
)
