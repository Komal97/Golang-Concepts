package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Book struct (Model)
type Book struct {
	ID string `json: "id, omitempty"`       // uuid
	Isbn string `json: "isbn,omitempty"`
	Title string `json: "title,omitempty"`
	Author *Author `json: "author,omitempty"`
}

// Author struct (Model)
type Author struct{
	FirstName string `json: "firstname,omitempty"`
	LastName string `json: "lastname,omitempty"`
}

// Get all books
func getBooks(resp http.ResponseWriter, req *http.Request)  {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(books)
}

// Get single book
func getBook(resp http.ResponseWriter, req *http.Request)  {
	resp.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)  // get path parameters

	fmt.Println("Params: ", params)
	for _, book := range books{
		if book.ID == params["id"]{
			json.NewEncoder(resp).Encode(book)
			return
		}
	}

	// if id doesn't exist then return empty book
	json.NewEncoder(resp).Encode(&Book{})
}

// Create new book
func createBook(resp http.ResponseWriter, req *http.Request)  {
	resp.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(req.Body).Decode(&book)
	// book.ID = strconv.Itoa(rand.Intn(1000000))
	book.ID = "3"
	books = append(books, book)
	fmt.Println(books)
	json.NewEncoder(resp).Encode(book)
}

// Update a book record
func updateBook(resp http.ResponseWriter, req *http.Request)  {

	resp.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)  // get path parameters
	fmt.Println("Params: ", params)

	for i, book := range books{
		if book.ID == params["id"]{
			books = append(books[:i], books[i+1:]...)
			var b Book
			_ = json.NewDecoder(req.Body).Decode(&b)
			b.ID = params["id"]
			books = append(books, b)
			break
		}
	}
	json.NewEncoder(resp).Encode(books)
}

// Delete a book record
func deleteBook(resp http.ResponseWriter, req *http.Request)  {
	resp.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)  // get path parameters
	fmt.Println("Params: ", params)

	for i, book := range books{
		if book.ID == params["id"]{
			books = append(books[:i], books[i+1:]...)
			break
		}
	}
	json.NewEncoder(resp).Encode(books)
}