package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Linabousbih/go-bookstore/pkg/models"
	"github.com/Linabousbih/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Conetnt-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	ID, err := strconv.ParseInt(param["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	newBook, _ := models.GetBookById(ID)
	res, _ := json.Marshal(newBook)
	w.Header().Set("Conetnt-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	deleted := models.DeleteBook(ID)
	res, _ := json.Marshal(deleted)
	w.Header().Set("Conetnt-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	oldBook, db := models.GetBookById(ID)
	if updatedBook.Name != "" {
		oldBook.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		oldBook.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		oldBook.Publication = updatedBook.Publication
	}
	db.Save(&oldBook)
	res, _ := json.Marshal(oldBook)
	w.Header().Set("Conetnt-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
