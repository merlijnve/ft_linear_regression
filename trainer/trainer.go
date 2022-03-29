package main

import (
	"fmt"
	"os"
	"strconv"
)

type Coordinate struct {
	X float64
	Y float64
}

type Dataset struct {
	xName string
	yName string
	xMin  float64
	xMax  float64
	yMin  float64
	yMax  float64
	data  []Coordinate
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func writeThetaValues(theta0 float64, theta1 float64) {
	file, err := os.Create("../thetaValues.txt")
	handleError(err, "Error: could not create thetaValues file")
	defer file.Close()

	text := "theta0:" + strconv.FormatFloat(theta0, 'f', -1, 64) + "\n" +
		"theta1:" + strconv.FormatFloat(theta1, 'f', -1, 64) + "\n"
	_, err = file.Write([]byte(text))
	handleError(err, "Error: could not write thetaValues to file")
	fmt.Printf("Wrote theta values to file: \"../thetaValues.txt\"\nTheta0: %f\nTheta1: %f\n",
		theta0, theta1)
}


func main() {
	dataset := readDataset()
	plotData(dataset, "data_raw.png")

	if len(os.Args) != 3 {
		fmt.Println("Error: no correct input\nUse: ./trainer [dataset filename] [learningRate]")
		os.Exit(0)
	}
	learningRate, err := strconv.ParseFloat(os.Args[2], 64)
	handleError(err, "Error: no correct learningRate given\nUse: ./trainer [dataset filename] [learningRate]")
	theta0, theta1 := gradientDescent(learningRate, dataset)
	writeThetaValues(theta0, theta1)
	plotDataAndLinearEquation(dataset, theta0, theta1)
}
