package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	/*Connect to MySQL db*/
	connStr := os.Getenv("MYSQL_CONN_STR")
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db, err)
}
