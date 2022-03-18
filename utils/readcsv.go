package utils

import (
	"encoding/csv"
	"fmt"
	"homework-3-ErdemOzgen/domain/book"
	"os"
	"strconv"
	"sync"
)

func ReadCsv(filename string) ([]book.Book, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	//reader.Comma = ';' // eğer virgül dışında farklı bir karakter ile verileri ayırıyorsanız
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result []book.Book

	for _, line := range lines[1:] {
		price, _ := strconv.Atoi(line[3])       // TODO: Handle error
		stockamount, _ := strconv.Atoi(line[4]) // TODO: Hadnle error
		data := book.Book{
			Name:              line[0],
			Author:            line[1],
			AuthorDescription: line[2],
			Price:             price,
			StockAmount:       stockamount,
		}

		result = append(result, data)
	}

	return result, nil
}

func ReadBOOKWithWorkerPool(path string) []book.Book {
	var resultBook []book.Book
	jobs := make(chan []string, 5)
	results := make(chan book.Book)

	wg := new(sync.WaitGroup)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go convertToBookStruct(jobs, results, wg)
	}

	go func() {
		f, _ := os.Open(path)
		defer f.Close()
		lines, _ := csv.NewReader(f).ReadAll()
		isFirstRow := true
		for _, line := range lines {
			if isFirstRow {
				isFirstRow = false
				continue
			}

			jobs <- line
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()

		close(results)
	}()

	for v := range results {
		fmt.Println(v)
		//bookRepository.InsertBook(v)
		resultBook = append(resultBook, v)
	}

	//return []book.Book, nil
	return resultBook
}

func convertToBookStruct(jobs <-chan []string, results chan<- book.Book, wg *sync.WaitGroup) {
	defer wg.Done()
	// eventually I want to have a []string channel to work on a chunk of lines not just one line of text
	for j := range jobs {
		price, _ := strconv.Atoi(j[3])       // TODO: Handle error
		stockamount, _ := strconv.Atoi(j[4]) // TODO: Handle error
		b := book.Book{
			Name:              j[0],
			Author:            j[1],
			AuthorDescription: j[2],
			Price:             price,
			StockAmount:       stockamount,
		}
		results <- b
	}
}
