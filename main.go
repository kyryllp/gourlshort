package main

import (
	"fmt"
	"os"
)

// no need for this if using docker
//func init() {
//	if err := godotenv.Load(); err != nil {
//		log.Fatalln("No .env file found")
//	}
//}

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	fmt.Println("Server is up and running...")
	a.Run(":8080")
}
