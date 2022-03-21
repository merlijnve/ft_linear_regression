package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"image/color"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

type ThetaValues struct {
	theta0 float64 `json:"theta0"`
	theta1 float64 `json:"theta1"`
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(0)
	}
}

func readThetaValues() (float64, float64) {
	var thetaValues ThetaValues
	
	// handle file open & close
	file, err := os.Open("../thetaValues.txt")
	handleError(err, "Error: could not read file")
	defer file.Close()

	// read theta0 and theta1 from file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	thetaValues.theta0, err = strconv.ParseFloat(strings.Split(scanner.Text(), ":")[1], 64)
	handleError(err, "Error: theta0 in file is not a number")
	scanner.Scan()
	thetaValues.theta1, err = strconv.ParseFloat(strings.Split(scanner.Text(), ":")[1], 64)
	handleError(err, "Error: theta1 in file is not a number")
	
	return thetaValues.theta0, thetaValues.theta1
}

func promptMileage() float64 {
	var mileagePointer *float64

	fmt.Println("\nPlease enter a mileage and press [enter]:")
	reader := bufio.NewReader(os.Stdin)
	for mileagePointer == nil {
		text, err := reader.ReadString('\n')
		handleError(err, "Error: could not read mileage")
		mileage, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
		if err != nil {
			fmt.Println("That's not a correct mileage, please try again:")
		} else {
			mileagePointer = &mileage
		}
	}
	return *mileagePointer
}

func plotLinearEquation(theta0 float64, theta1 float64, mileage float64) {
	p := plot.New()

	p.Title.Text = "Car price prediction"
	p.X.Label.Text = "Mileage"
	p.Y.Label.Text = "Price "

	linearEquation := plotter.NewFunction(func(x float64) float64 {
			return theta0 + (theta1 * x)
		})
	linearEquation.Color = color.RGBA{B: 255, A: 255}

	p.Add(linearEquation)
	p.X.Min = 0
	p.X.Max = 100
	p.Y.Min = 0
	p.Y.Max = 100

	err := p.Save(600, 600, "functions.png")
	handleError(err, "Error: could not save picture of linear equation")
}

func main() {
	fmt.Println("Car price predictor")
	fmt.Println("--------------------")
	
	theta0, theta1 := readThetaValues()
	mileage := promptMileage()
	estimatePrice := theta0 + (theta1 * mileage)
	fmt.Printf("Estimated price with mileage %s:\n", strconv.FormatFloat(mileage, 'f', -1, 64))
	fmt.Println(estimatePrice)
	plotLinearEquation(theta0, theta1, mileage)
}