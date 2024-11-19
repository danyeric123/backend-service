package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DB is the database connection
var DB *sqlx.DB

// Connect connects to the database
func Connect() {
	var err error
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
	)
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to the database")
	}
	logrus.Info("Connected to the database")
}
