package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_DSN                       string
	DBUser                       string
	DBPassword                   string
	DBName                       string
	InstanceConnectionName       string
	Mode                         string
	GoogleApplicationCredentials string
	GoogleAccessID               string
	Name                         string
	Basket                       string
	JWTSecret                    string
)

// .envを呼び出します。
func LoadEnv() {
	if os.Getenv("MODE") == "dev" {
		err := godotenv.Load()

		if err != nil {
			log.Printf("読み込み出来ませんでした: %v", err)
		}
	}

	DB_DSN = os.Getenv("DB_DSN")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	InstanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
	Mode = os.Getenv("MODE")
	// GoogleApplicationCredentials = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	// GoogleAccessID = os.Getenv("GOOGLE_ACCESS_ID")
	// Name = os.Getenv("NAME")
	Basket = os.Getenv("GCP_BUCKET_ID")
	JWTSecret = os.Getenv("JWT_SECRET")
}
