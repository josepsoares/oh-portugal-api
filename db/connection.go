package db

import (
	"fmt"
	"josepsoares/iberiapi/utils"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBConn *gorm.DB
)

func Connect() {
	err := godotenv.Load()
	utils.FailOnError("couldn't read .env file", err)

	var (
		pg_user     = os.Getenv("POSTGRES_USER")
		pg_password = os.Getenv("POSTGRES_PASSWORD")
		pg_host     = os.Getenv("POSTGRES_HOST")
		pg_db       = os.Getenv("POSTGRES_DB")
		pg_port     = os.Getenv("POSTGRES_PORT")
	)

	port, err := strconv.Atoi(pg_port)
	utils.FailOnError("couldn't parse port", err)

	dsn := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", pg_user, pg_password, pg_host, pg_db, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")

	DBConn = db
}
