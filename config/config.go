package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	// Server
	AppPort string
	AppEnv  string
	AppKey  string

	// Database
	DBDriver   string
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUsername string
	DBPassword string

	// JWT
	JWTSecret          string
	JWTExpirationHours int

	// Mail
	MailHost     string
	MailPort     string
	MailUsername string
	MailPassword string
	MailFrom     string
	MailFromName string

	// File Storage
	StoragePath string
	UploadPath  string
}

var AppConfig Config

// Load loads configuration from .env file and environment variables
func Load() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	AppConfig = Config{
		AppPort: getEnv("APP_PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "production"),
		AppKey:  getEnv("APP_KEY", "base64:change_me_to_a_secure_random_key"),

		DBDriver:   getEnv("DB_CONNECTION", "mysql"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBDatabase: getEnv("DB_DATABASE", "kms"),
		DBUsername: getEnv("DB_USERNAME", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),

		JWTSecret:          getEnv("JWT_SECRET", "change_me_to_a_secure_jwt_secret"),
		JWTExpirationHours: 24,

		MailHost:     getEnv("MAIL_HOST", "smtp.mailtrap.io"),
		MailPort:     getEnv("MAIL_PORT", "2525"),
		MailUsername: getEnv("MAIL_USERNAME", ""),
		MailPassword: getEnv("MAIL_PASSWORD", ""),
		MailFrom:     getEnv("MAIL_FROM_ADDRESS", "noreply@kms.local"),
		MailFromName: getEnv("MAIL_FROM_NAME", "KMS"),

		StoragePath: getEnv("STORAGE_PATH", "./storage"),
		UploadPath:  getEnv("UPLOAD_PATH", "./public/uploads"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
