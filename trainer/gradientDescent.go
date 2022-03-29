package main

import (
	"math"
)

// linear equation for estimating price
func estimatePrice(mileage float64, theta0 float64, theta1 float64) float64 {
	return theta0 + (theta1 * mileage)
}

func normalizeDataset(dataset Dataset) Dataset {
	var normalizedDataset Dataset

	// create a copy of the dataset
	normalizedDataset.xName = dataset.xName
	normalizedDataset.yName = dataset.yName
	normalizedDataset.xMax = dataset.xMax
	normalizedDataset.xMin = dataset.xMin
	normalizedDataset.yMax = dataset.yMax
	normalizedDataset.yMin = dataset.yMin
	normalizedDataset.data = make([]Coordinate, len(dataset.data))
	copy(normalizedDataset.data, dataset.data)

	// normalize
	for i, c := range dataset.data {
		normalizedDataset.data[i].Y = (c.Y - dataset.yMin) / (dataset.yMax - dataset.yMin)
		normalizedDataset.data[i].X = (c.X - dataset.xMin) / (dataset.xMax - dataset.xMin)
	}
	
	return normalizedDataset
}

func denormalizeThetas(dataset Dataset, theta0 float64, theta1 float64) (float64, float64) {

	theta1 = (dataset.yMax - dataset.yMin) * theta1 / (dataset.xMax - dataset.xMin)
	theta0 = dataset.yMin + ((dataset.yMax - dataset.yMin) * theta0) + theta1 * (1 - dataset.xMin)
	
	return theta0, theta1
}

func gradientDescent(learningRate float64, dataset Dataset) (float64, float64) {
	var theta0, theta1 float64 = 0, 0
	var delta0, delta1 float64 = 1, 1
	var prevCost0, prevCost1 float64
	normalizedDataset := normalizeDataset(dataset)
	data := normalizedDataset.data
	m := float64(len(data))
	

	for delta0 > 0.000001 || delta1 > 0.000001 {
		cost0 := 0.0
		cost1 := 0.0
		for _, c := range data {
			cost0 += estimatePrice(c.X, theta0, theta1) - c.Y
			cost1 += (estimatePrice(c.X, theta0, theta1) - c.Y) * c.X
		}
		delta0 = math.Abs(prevCost0 - cost0)
		delta1 = math.Abs(prevCost1 - cost1)
		theta0 -= learningRate * (cost0 / m)
		theta1 -= learningRate * (cost1 / m)
		prevCost0 = cost0
		prevCost1 = cost1
	}


	theta0, theta1 = denormalizeThetas(dataset, theta0, theta1) 

	return theta0, theta1
}
