package book

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	result := r.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (r *BookRepository) GetBookByID(id uint) (Book, error) {
	var book Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

func (r *BookRepository) GetBookByName(name string) (Book, error) {
	var book Book
	result := r.db.Where(Book{Name: name}).Attrs(Book{Name: "NULL"}).FirstOrInit(&book)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

func (r *BookRepository) GetBookByNameOrCreate(name string) (Book, error) {
	var book Book
	result := r.db.Where(Book{Name: name}).Attrs(Book{Name: name}).FirstOrCreate(&book)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

func (r *BookRepository) Migration() {
	r.db.AutoMigrate(&Book{})
}

func (r *BookRepository) FindBookByName(name string) (Book, error) {
	var book Book
	result := r.db.Where("name LIKE ?", "%"+name+"%").First(&book)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

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

func (r *BookRepository) InsertSampleDataFromSlices(s []Book) {
	for _, b := range s {
		r.db.Create(&b)
	}
}

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

func (r *BookRepository) DeleteBookByID(id int) error {
	var book Book
	result := r.db.First(&book, id)
	//find book with id and delete it
	//result := r.db.First(&book, id).Where(Book{IsDelete: false})
	//r.db.Raw("Select * from book where id = ?", id).Scan(&book)
	//fmt.Println("Book deleted =====> ", book)
	//if book.IsDelete {
	//	fmt.Println("ALREADY DELETED")
	//	return nil
	//}
	//result := r.db.Unscoped().Where(&Book.{gorm.Model: gorm.Model{ID: id}}).Delete(&book)
	//result := r.db.First(&book, id)
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

// ADD ISDELETD
func (r *BookRepository) DeleteBookByID2(id int) error {
	var book Book
	result := r.db.Unscoped().Where("ID=?", id).Find(&book)
	if result.Error != nil {
		log.Println("error: ", result.Error)
		return result.Error
	}
	fmt.Println("adjsnajdnsanjdnasnkjdnnjk")
	fmt.Println(book)
	/*book.IsDelete = true
	if book.IsDelete {
		fmt.Println("ALREADY DELETED")
		return nil

	}*/
	fmt.Println("Book has been deleted...", book)
	result = r.db.Delete(&book)
	if result.Error != nil {
		log.Println("error: ", result.Error)
		return result.Error
	}
	return nil
}

func (r *BookRepository) InsertBook(book *Book) {
	r.db.Create(&book)
}
