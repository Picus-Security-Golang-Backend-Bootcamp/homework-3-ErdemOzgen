package main

import (
	"fmt"
	"gorm-tut/domain/book"
	"gorm-tut/domain/city"
	"gorm-tut/domain/country"
	"gorm-tut/domain/erdem"
	"gorm-tut/infrastructure"
	"gorm-tut/utils"
)

var (
	repository        *city.CityRepository
	countryRepository *country.CountryRepository
	erdemRepo         *erdem.ErdemRepository
	bookRepository    *book.BookRepository
)

func init() {
	db := infrastructure.NewMySQLDB("root:root@tcp(127.0.0.1:3306)/library?parseTime=True&loc=Local")
	repository = city.NewCityRepository(db)
	countryRepository = country.NewCountryRepository(db)
	repository.Migration()
	repository.InsertSampleData()
	countryRepository.Migration()
	countryRepository.InsertSampleData()
	erdemRepo = erdem.NewErdemRepository(db)
	erdemRepo.Migration()
	erdemRepo.InsertSampleData()
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
