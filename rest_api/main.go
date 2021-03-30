package main

import(
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var books []Book

func main(){

	// Initialize router
	router := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "847564", Title: "Book Two", Author: &Author{FirstName: "Jane", LastName: "Smith"}})

	// Route Handlers for establishing end-points
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run server
	err := http.ListenAndServe("localhost:8000", router)
	if err != nil{
		fmt.Println("Error while starting server: ", err)

	} else{
		fmt.Println("Server started")
	}
}

