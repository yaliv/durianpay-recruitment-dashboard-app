package config

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

var (
	JwtSecret           = []byte(getEnv("JWT_SECRET", "dev-secret-replace-me"))
	JwtExpired          = getEnv("JWT_EXPIRED", "24h")
	HttpAddress         = getEnv("HTTP_ADDR", ":8080")
	Port                = getEnv("PORT", "8080")
	CorsAllowedOrigins  = getEnv("CORS_ALLOWED_ORIGINS", "")
	OpenapiYamlLocation = getEnv("OPENAPIYAML_LOCATION", "../openapi.yaml")
)

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
