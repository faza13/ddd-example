package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func NewConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	App.Env = os.Getenv("APP_ENV")
	App.Port = os.Getenv("APP_PORT")
	App.Name = os.Getenv("APP_NAME")
	App.Timezone = os.Getenv("APP_TIMEZONE")

	Database.Host = os.Getenv("DB_HOST")
	Database.Port = os.Getenv("DB_PORT")
	Database.Database = os.Getenv("DB_DATABASE")
	Database.Username = os.Getenv("DB_USERNAME")
	Database.Password = os.Getenv("DB_PASSWORD")
}
