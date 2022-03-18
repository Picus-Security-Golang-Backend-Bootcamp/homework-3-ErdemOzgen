package book

import (
	"fmt"

	"gorm.io/gorm"
)

/*
// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
*/

type Book struct {
	gorm.Model
	Name              string
	Author            string
	AuthorDescription string
	Price             int //TODO:	float64
	StockAmount       int
	//IsDelete          bool
}

func (b *Book) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("City (%s) deleting...", b.Name)
	return nil
}
