package book

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

//Function returns gorm.DB instance
type BookRepository struct {
	db *gorm.DB
}

//functon creates a new BookRepository
func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

//Gets all books
func (r *BookRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	result := r.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// Gets book by id and returns book and errors
func (r *BookRepository) GetBookByID(id uint) (Book, error) {
	var book Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

// Get function by name and returns book
// TODO: returns slice not only one book
func (r *BookRepository) GetBookByName(name string) (Book, error) {
	var book Book
	result := r.db.Where(Book{Name: name}).Attrs(Book{Name: "NULL"}).FirstOrInit(&book)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

// Try to get book by name and returns book and error if not found it will createed new in DB
func (r *BookRepository) GetBookByNameOrCreate(name string) (Book, error) {
	var book Book
	result := r.db.Where(Book{Name: name}).Attrs(Book{Name: name}).FirstOrCreate(&book)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

//Function for handinling migration
func (r *BookRepository) Migration() {
	r.db.AutoMigrate(&Book{})
}

// function takes book name and returns book and error
func (r *BookRepository) FindBookByName(name string) (Book, error) {
	var book Book
	result := r.db.Where("name LIKE ?", "%"+name+"%").First(&book)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

// Inserts sample data test values
func (r *BookRepository) InsertSampleData() {
	book := []Book{
		{Name: "Book1", Author: "Author1", AuthorDescription: "AuthorDescription1", Price: 100, StockAmount: 10},
		{Name: "Book2", Author: "Author2", AuthorDescription: "AuthorDescription2", Price: 200, StockAmount: 20},
		{Name: "Book3", Author: "Author3", AuthorDescription: "AuthorDescription3", Price: 300, StockAmount: 30},
		{Name: "Book4", Author: "Author4", AuthorDescription: "AuthorDescription4", Price: 300, StockAmount: 30},
		{Name: "Book5", Author: "Author5", AuthorDescription: "AuthorDescription5", Price: 300, StockAmount: 30},
	}
	for _, b := range book {
		r.db.Create(&b)
	}

}

// inserts sample data from slice
func (r *BookRepository) InsertSampleDataFromSlices(s []Book) {
	for _, b := range s {
		r.db.Create(&b)
	}
}

// Updates book from book structs
func (r *BookRepository) UpdateBook(b Book, id int) error {
	//r.db.First(&book, id)
	var book Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return result.Error
	}
	fmt.Println(book)
	//change Fields
	book.Name = b.Name
	book.Author = b.Author
	book.AuthorDescription = b.AuthorDescription
	book.Price = b.Price
	book.StockAmount = b.StockAmount
	//save
	result = r.db.Save(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Buys book from id and returns  error
func (r *BookRepository) BuyBookByID(id int, amount int) error {
	var book Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return result.Error
	}
	if book.StockAmount < amount {
		fmt.Println("NOT ENOUGH STOCK")
		return nil
	}
	book.StockAmount -= amount
	result = r.db.Save(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//Delete book by IDs and returns error
func (r *BookRepository) DeleteBookByID(id int) error {
	var book Book
	result := r.db.First(&book, id)

	if result.Error != nil {
		log.Println("error: NOT FOUND ALREADY DELETED ", result.Error)

		return result.Error
	}

	log.Println(book)
	result = r.db.Delete(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

//Insert book into database takes book struct
func (r *BookRepository) InsertBook(book *Book) {
	r.db.Create(&book)
}
