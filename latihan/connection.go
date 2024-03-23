package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)
func main(){
	// Load .env file
    if err := godotenv.Load(); err != nil {
        fmt.Println("Error loading .env file:", err)
        return
    }
	
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/db_golang", username, password))
	// db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:3306)/db_golang", username))


	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
		
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
        fmt.Println("Error pinging database:", err)
        return
    }

    fmt.Println("Connected to database successfully")

}
