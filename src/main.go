package main

import (
	"exercise-battleship-2/src/model"
	"fmt"
)

func main() {
	var filepath string
	fmt.Println("Enter file path: ")
	fmt.Scan(&filepath)

	file := model.NewFileReader(filepath)

}
