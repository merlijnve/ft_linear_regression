package main

import (
	"fmt"
	"os"
)

type Coordinate struct {
	X float64
	Y float64
}

type Dataset struct {
	xName string
	yName string
	data []Coordinate
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func main() {
	dataset := readDataset()
	plotData(dataset)
}
