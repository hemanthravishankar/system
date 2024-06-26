package system

import (
	"github.com/hemanthravishankar/elements"
)

// Implement the following functionalities:
var books []elements.Book
var members []elements.Member

// var loan []elements.Loan
var bookID, memberID int //loanID

// Add Book: Add a new book to the library. ID, Title, Author, YearPublished, IsAvailable.
func AddBook(title string, author string, yearpub int) elements.Book {
	bookID++
	isavl := true
	book := elements.Book{ID: bookID, Title: title, Author: author, YearPublished: yearpub, IsAvailable: isavl}
	books = append(books, book)
	return book
}

// Register Member: Register a new member to the library. ID, Name, Email.
func AddMember(name string, email string) elements.Member {
	memberID++
	member := elements.Member{ID: memberID, Name: name, Email: email}
	members = append(members, member)
	return member

}

//Loan Book: Loan a book to a member. Ensure that a book can only be loaned if it is available.

//Return Book: Return a book to the library.

// List Available Books: List all available books in the library.
func Getbooks() []elements.Book {
	return books
}

// List Members: List all registered members.
func GetMembers() []elements.Member {
	return members
}

//List Loans: List all current loan records.
