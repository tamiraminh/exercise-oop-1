package main

import (
	"exercise-battleship-2/src/model/battleship"
	"exercise-battleship-2/src/model/file"
	"fmt"
)

func main() {
	var filepath string
	fmt.Println("Enter file path: ")
	fmt.Scan(&filepath)

	file := file.NewFileReader()
	file.ReadFile(filepath)

	game := battleship.NewGame(file.GetNumShip(), file.GetGridSize(), file.GetNumMissile(), file.GetPlayer1Ship(), file.GetPlayer2Ship())
	game.ProceedGame(file.GetPlayer1Missile(), file.GetPlayer2Missile())
	game.Result()

}
