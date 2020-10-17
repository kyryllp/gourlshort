package db

import (
	"database/sql"
	"fmt"
	"gourlshort/model"
	"log"
)

func InitializeConnection(user string, password string, dbname string) (*sql.DB, error) {
	//dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", user, password, dbname)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(db:3306)/%s", user, password, dbname)
	return sql.Open("mysql", dataSourceName)
}

func GetUrl(db *sql.DB, name string) model.URL {
	result, err := db.Query("SELECT id, redirect_name, original_url FROM urls WHERE redirect_name = ?", name)
	if err != nil {
		log.Fatalln(err)
	}

	defer result.Close()

	var url model.URL
	for result.Next() {
		err := result.Scan(&url.ID, &url.RedirectName, &url.OriginalUrl)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return url
}

func SaveUrl(db *sql.DB, url model.URL) (sql.Result, error) {
	stmt, err := db.Prepare("INSERT INTO urls(redirect_name, original_url) VALUES(?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	return stmt.Exec(url.RedirectName, url.OriginalUrl)
}
