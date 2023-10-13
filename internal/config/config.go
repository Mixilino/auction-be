package config

type AppConfig struct {
	DBHost     string `env:"DB_HOST"`
	DBPort     int    `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBDriver   string `env:"DB_DRIVER"`
	DBName     string `env:"DB_NAME"`
	SSLMode    string `env:"SSL_MODE"`
	JWTSecret  string `env:"JWT_SECRET"`
}
