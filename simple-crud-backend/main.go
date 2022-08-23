package simplecrudbackend

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Note struct {
	Id      int    `gorm:"primary_key;column:id;"`
	Title   string `json:"title" gorm:"not null;column:title;"`
	Content string `json:"content" gorm:"not null;column:content;"`
}

type NoteUpdate struct {
	Title   *string
	Content *string
}

func (Note) TableName() string {
	return "notes"
}

func Main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	/// CREATE
	note := Note{
		Title:   "This is a title",
		Content: "This is a content",
	}
	db.Create(&note)

	fmt.Println("Start reading")
	time.Sleep(time.Second * 2)

	/// READ
	readNote := Note{}
	db.Where("title = ?", "This is a title").First(&readNote)
	fmt.Printf("Read note : %+v\n", readNote)

	fmt.Println("Start updating")
	time.Sleep(time.Second * 2)

	/// UPDATE
	newTitle, newContent := "This is a new title", ""
	db. //
		Table(Note{}.TableName()).
		Where("id = ?", readNote.Id).
		Updates(NoteUpdate{
			Title:   &newTitle,
			Content: &newContent,
		})
	db.Where("id = ?", readNote.Id).First(&readNote)
	fmt.Printf("Read note after update : %+v\n", readNote)

	fmt.Println("Start deleting")
	time.Sleep(time.Second * 2)

	/// DELETE
	db.Table(Note{}.TableName()).Where("id = ?", readNote.Id).Delete(nil)
}
