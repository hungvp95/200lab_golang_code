package simplecrudbackend

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Note struct {
	Id      int
	Title   string
	Content string
}

func Main() {
	dsn := os.Getenv("DBConnectionStr")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	note := Note{
		Title: "Hello World",
		Content: "Getting started with golang",
	}
	db.Create(&note)
	
	fmt.Printf("%+v\n", note)
}
