package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintPretty(v interface{}) {
	vstruct, err := json.MarshalIndent(v, " ", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(vstruct))

}
