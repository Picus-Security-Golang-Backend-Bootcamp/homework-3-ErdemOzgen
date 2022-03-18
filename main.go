package main

import (
	"fmt"
	"homework-3-ErdemOzgen/domain/book"
	"os"
	"strconv"
	"strings"

	"homework-3-ErdemOzgen/infrastructure"
	"homework-3-ErdemOzgen/utils"
)

var (
	bookRepository *book.BookRepository
)

func init() {
	db := infrastructure.NewMySQLDB("root:root@tcp(127.0.0.1:3306)/library?parseTime=True&loc=Local")
	//====================>
	IsInitExecuted := utils.CheckInitExecuted("hasinitexecuted.bool")
	bookRepository = book.NewBookRepository(db)
	bookRepository.Migration()

	if !IsInitExecuted {
		bookRepository.InsertSampleData()
		er := utils.ReadBOOKWithWorkerPool("books.csv")
		bookRepository.InsertSampleDataFromSlices(er)
	}

}

func main() {
	/*
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
		fmt.Println("===================")
		// read ReadCsv
		//er := utils.ReadBOOKWithWorkerPool("books.csv")
		//utils.PrintPretty(er)
	*/
	if len(os.Args) == 1 {
		utils.PrintHelp()
		return
	}

	cmds := os.Args[1]

	switch strings.ToLower(cmds) {
	case "buy", "b", "-b", "--buy":
		fmt.Println("Buying...")
		if len(os.Args) <= 2 {
			fmt.Println("Not enough args for buy operation")
			return
		}
		i, _ := strconv.Atoi(os.Args[2])

		j, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(utils.ErrNotInt)
			return
		}
		err = bookRepository.BuyBookByID(i, j)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Buy function has been executed")

	case "delete", "d", "--delete", "-d":
		fmt.Println("Deleting...")

		if len(os.Args) <= 2 {
			fmt.Println("Not enough args for delete operation")
			return
		}
		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(utils.ErrNotInt)
			return
		}
		err = bookRepository.DeleteBookByID(i)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Delete function has been executed")

	case "search", "s", "--search", "-s":
		fmt.Println("Searching...")
		if len(os.Args) <= 2 {
			fmt.Println("Not enough args for search operation")
			return
		}
		e, err := bookRepository.FindBookByName(os.Args[2])
		if err != nil {
			fmt.Println(utils.ErrNotFound)
			return
		}
		utils.PrintPretty(e)

	case "list", "l", "--list", "-l":
		fmt.Println("Listing...")
		listedBooks, _ := bookRepository.GetAllBooks()
		utils.PrintPretty(listedBooks)

	case "help", "h", "--help", "-h":
		utils.PrintHelp()
	case "version", "v", "--version", "-v":
		utils.PrintBanner()

	default:
		fmt.Println("Unknown command")

	}
}
