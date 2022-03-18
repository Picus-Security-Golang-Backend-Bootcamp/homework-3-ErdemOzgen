package utils

import (
	"bufio"
	"log"
	"os"
)

func CheckInitExecuted(fileName string) bool {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		//fmt.Printf("line: %s\n", scanner.Text())
		if scanner.Text() == "1" {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return false
}
