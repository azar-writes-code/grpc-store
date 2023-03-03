package configs

import (
	"os"

	"github.com/joho/godotenv"
)

func GetMongoURL() string {
	err := godotenv.Load()
	if err != nil {
		return ""
	}

	return os.Getenv("MONGODB_URL")

}
