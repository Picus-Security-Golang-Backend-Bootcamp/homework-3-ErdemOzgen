package main

import (
	"fmt"
	"homework-3-ErdemOzgen/domain/book"

	"homework-3-ErdemOzgen/infrastructure"
	"homework-3-ErdemOzgen/utils"
)

var (
	bookRepository *book.BookRepository
)

func init() {
	db := infrastructure.NewMySQLDB("root:root@tcp(127.0.0.1:3306)/library?parseTime=True&loc=Local")
	//====================>
	bookRepository = book.NewBookRepository(db)
	bookRepository.Migration()
	bookRepository.InsertSampleData()

}

func main() {
	fmt.Println("Get All Book")
	listedBooks, _ := bookRepository.GetAllBooks()
	utils.PrintPretty(listedBooks)
	fmt.Println("===================")
	fmt.Println("Buy Book By ID")
	//bookRepository.BuyBookByID(2, 5)
	fmt.Println("================")
	fmt.Println("Delete By ID")
	//bookRepository.DeleteBookByID(4)
	fmt.Println("===================")
	fmt.Println("Get Book by ID")
	e, _ := bookRepository.GetBookByID(1)
	utils.PrintPretty(e)
	fmt.Println("===================")
	fmt.Println("Find Book by Name")
	e, _ = bookRepository.FindBookByName("Book1")
	utils.PrintPretty(e)
	fmt.Println("===================")
	fmt.Println("Update Book ")
	bookRepository.UpdateBook(book.Book{Name: "Erdem", Author: "Author1", AuthorDescription: "AuthorDescription1", Price: 100, StockAmount: 10}, 1)
	fmt.Println("===================")
	bookRepository.DeleteBookByID(1)
}
