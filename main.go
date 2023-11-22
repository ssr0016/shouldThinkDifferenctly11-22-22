package main

import "time"

type Book struct {
	title    string
	author   string
	numPages int

	isSaved bool
	savedAT time.Time
}

// saved Book, same function as below saveBook function
func (book *Book) saveBook() {
	book.isSaved = true
	book.savedAT = time.Now()
}

// saved Book
func saveBook(book *Book) {
	book.isSaved = true
	book.savedAT = time.Now()
}

// 1. read data
// 2. write data

func main() {

}

// Note go lang is all about data, structures, data and functions
