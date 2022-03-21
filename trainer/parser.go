package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

func splitLine(line string) (string, string) {
	values := strings.Split(line, ",")
	return values[0], values[1]
}

func splitLineToCoordinate(line string) Coordinate {
	values := strings.Split(line, ",")

	f1, err := strconv.ParseFloat(values[0], 64)
	handleError(err, "Error: non-number found in input file")
	f2, err2 := strconv.ParseFloat(values[1], 64)
	handleError(err2, "Error: non-number found in input file")
	
	return Coordinate{f1, f2}
}

func readDataset() Dataset {
	var dataset Dataset

	// handle file open & close
	filename := "../data.csv"
	if (len(os.Args) > 1) {
		filename = os.Args[1]
	}
	file, err := os.Open(filename)
	handleError(err, "Error: could not read file")
	defer file.Close()

	// read values from file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	dataset.xName, dataset.yName = splitLine(scanner.Text())
	for scanner.Scan() {
		dataset.data = append(dataset.data, splitLineToCoordinate(scanner.Text()))
	}
	return dataset
}