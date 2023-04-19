package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Book struct {
	Title       string
	Author      string
	Genre       string
	PublishYear int
}

var db []Book

func init() {
	content, err := ioutil.ReadFile("./db/book-storage.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	json.Unmarshal([]byte(content), &db)
}

func (book *Book) CreateBook() *Book {
	return book
}

func GetAllBooks() []Book {
	return db
}
