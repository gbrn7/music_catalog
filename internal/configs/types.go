package configs

type (
	Config struct {
		Service       Service
		Database      Database
		SpotifyConfig SpotifyConfig
	}

	Service struct {
		Port      string
		SecretKey string
	}

	Database struct {
		DataSourceName string
	}

	SpotifyConfig struct {
		ClientID     string
		ClientSecret string
	}
)
