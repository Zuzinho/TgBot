package env

import (
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Token      string
	ConnConfig pgx.ConnConfig
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	var exists bool
	var err error

	Token, exists = os.LookupEnv("TG_BOT_TOKEN")
	if !exists {
		log.Println("TgBot token not found")
	}

	connString, exists := os.LookupEnv("DB_CONN_STRING")
	if !exists {
		log.Fatal("Connection string not found")
	}
	ConnConfig, err = pgx.ParseConnectionString(connString)
	if err != nil {
		log.Fatal(err)
	}
}
