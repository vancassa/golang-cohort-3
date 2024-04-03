package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var PORT = ":9090"

type Book struct {
	ID	int
	Title	string
	Stock	int
	Author string
}

var books = []Book {
	{ID: 1, Title: "Title book 1", Stock: 20, Author: "Author A"},
	{ID: 2, Title: "Title book 2", Stock: 10, Author: "Author B"},
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/book", createBook)

	// Menjalankan server aplikasi. Dua param: Port yg dipakai & http.Handler
	http.ListenAndServe(PORT, nil)
}

// http.ResponseWriter: interface dgn berbagai method untuk mengirim response balik ke client
// http.Request: struct yg digunakan untuk mendapat data request e.g. headers, form value, url param
func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello"
	fmt.Fprintf(w, msg)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// get books
		tpl, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, books)
		
		// json.NewEncoder(w).Encode(books)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		title := r.FormValue("title")
		author := r.FormValue("author")
		stock := r.FormValue("stock")
		convertStock, err := strconv.Atoi(stock)
		if err != nil {
			http.Error(w, "Invalid stock", http.StatusBadRequest)
		}

		newBook := Book{ ID: len(books)+1, Title: title, Author: author, Stock: convertStock}
		books = append(books, newBook)

		json.NewEncoder(w).Encode(books)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}


func getBook(w http.ResponseWriter, r *http.Request) {
}