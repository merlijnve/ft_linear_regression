package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type ThetaValues struct {
	theta0 int
	theta1 int
}

func handleError(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}

func readThetaValues() ThetaValues {
	var thetaValues ThetaValues
	
	jsonFile, err := os.Open("../thetaValues.json")
	if err != nil {
		handleError("Error: could not read file")
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &thetaValues)
	return thetaValues
}

func main() {
	var thetaValues = readThetaValues()
	fmt.Println(thetaValues.theta0, thetaValues.theta1)
}