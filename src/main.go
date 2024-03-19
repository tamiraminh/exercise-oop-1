package main

import (
	"exercise-battleship-2/src/model"
	"fmt"
)

func main() {
	var filepath string
	fmt.Println("Enter file path: ")
	fmt.Scan(&filepath)

	file := model.NewFileReader()
	file.ReadFile(filepath)

	game := model.NewGame(file.GetNumShip(), file.GetGridSize(), file.GetNumMissile(), file.GetPlayer1Ship(), file.GetPlayer2Ship())
	game.ProceedGame(file.GetPlayer1Missile(), file.GetPlayer2Missile())
	game.Result()

}
