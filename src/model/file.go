package model

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type FileReader struct {
	data           []string
	gridSize       int
	numShip        int
	player1ship    []string
	player2ship    []string
	numMissile     int
	player1Missile []string
	player2Missile []string
}

func NewFileReader(filepath string) *FileReader {
	f := FileReader{}

	return &f
}

func (f *FileReader) ReadFile(filepath string) {
	fmt.Printf("\n\nReading a file in Go lang\n")

	// file is read using ReadFile()
	// method of ioutil package
	data, err := ioutil.ReadFile(filepath)

	// in case of an error the error
	// statement is printed and
	// program is stopped
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", filepath)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

	f.data = strings.Split(string(data), "\n")

	// process data
	gridSizeint, _ := strconv.Atoi(f.data[0])
	numShipint, _ := strconv.Atoi(f.data[1])
	f.gridSize = gridSizeint
	f.numShip = numShipint

	f.player1ship = strings.Split(f.data[2], ":")
	f.player2ship = strings.Split(f.data[3], ":")
	numMissile, _ := strconv.Atoi(f.data[4])
	f.numMissile = numMissile
	f.player1Missile = strings.Split(f.data[5], ":")
	f.player2Missile = strings.Split(f.data[6], ":")
}

func (f FileReader) GetGridSize() int {
	return f.gridSize
}

func (f FileReader) GetNumShip() int {
	return f.numShip
}

func (f FileReader) GetNumMissile() int {
	return f.numMissile
}

func (f FileReader) GetPlayer1Ship() []string {
	return f.player1ship
}

func (f FileReader) GetPlayer2Ship() []string {
	return f.player2ship
}

func (f FileReader) GetPlayer1Missile() []string {
	return f.player1Missile
}

func (f FileReader) GetPlayer2Missile() []string {
	return f.player2Missile
}
