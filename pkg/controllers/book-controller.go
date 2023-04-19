package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/minhluannguyen/go-book-server/pkg/models"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)

	fmt.Fprintf(w, "%s", res)
}
