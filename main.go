package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}

func main() {
	a := App{}
	fmt.Println(os.Getenv("APP_DB_USERNAME"))
	fmt.Println(os.Getenv("APP_DB_PASSWORD"))
	fmt.Println(os.Getenv("APP_DB_NAME"))
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
