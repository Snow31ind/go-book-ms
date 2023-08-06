package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Snow31ind/go-book-ms/pkg/models"
	"github.com/Snow31ind/go-book-ms/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	b, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book, _ := models.GetBookById(id)
	b, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	if err := utils.ParseBody(r, book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p := book.Create()
	b, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	deletedBook := models.DeleteBook(id)
	b, err := json.Marshal(deletedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var book = &models.Book{}
	err := utils.ParseBody(r, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	currentBook, db := models.GetBookById(id)
	if book.Name != "" {
		currentBook.Name = book.Name
	}
	if book.Publication != "" {
		currentBook.Publication = book.Publication
	}
	if book.Author != "" {
		currentBook.Author = book.Author
	}
	db.Save(currentBook)
	b, err := json.Marshal(currentBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
