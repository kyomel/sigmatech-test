package main

import (
	"fmt"
	"log"
	"os"
	databases "sigmatech-test/pkg/db"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var (
	dbRepoConn databases.DatabaseRepo = databases.NewPostgres()
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	e := echo.New()

	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	portDB, _ := strconv.Atoi(dbPort)
	portConnect, _ := strconv.Atoi(port)

	db, err := dbRepoConn.Connect(dbHost, portDB, dbUser, dbPassword, dbName, sslMode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", portConnect)))
}
