package main

import (
	// DATABASE
	"database/sql"
	"final-project/database"
	"final-project/service"
	"fmt"

	_ "github.com/lib/pq"

	// ENV

	"github.com/joho/godotenv"
)

var (
	db  *sql.DB
	err error
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "071100191005Novi"
	dbname   = "food-order"
)

func main() {
	// Test load environment
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed to load environment.")
	} else {
		fmt.Println("Succesfully load environment.")
	}
	// Connecting Database
	// railwayInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("DB Connection Failed.")
		panic(err)
	} else {
		fmt.Println("DB Connection Success.")
	}

	err = db.Ping()
	service.CheckError(err)

	fmt.Println("Successfully connected to database")
	database.DbMigrate(db)
	defer db.Close()

}
