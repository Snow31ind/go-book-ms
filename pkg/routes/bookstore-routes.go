package routes

import (
	"github.com/Snow31ind/go-book-ms/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/books/{bookId}", controllers.UpdateBookById).Methods("PUT")
	r.HandleFunc("/books/{bookId}", controllers.DeleteBookById).Methods("DELETE")
}
