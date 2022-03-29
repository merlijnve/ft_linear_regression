package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

func coordsToXYs(data []Coordinate) plotter.XYs {
	pts := make(plotter.XYs, len(data))

	for i := range data {
		pts[i].X = data[i].X
		pts[i].Y = data[i].Y
	}

	return pts
}

func getMaxValues(data []Coordinate) (float64, float64) {
	var xMax float64 = 0
	var yMax float64 = 0

	for i := range data {
		if data[i].X > xMax {
			xMax = data[i].X
		}
		if data[i].Y > yMax {
			yMax = data[i].Y
		}
	}

	return xMax, yMax
}

func plotDataAndLinearEquation(dataset Dataset, theta0 float64, theta1 float64) {
	p := plot.New()
	
	p.Title.Text = "Car price prediction"
	p.X.Label.Text = dataset.xName
	p.Y.Label.Text = dataset.yName

	linearEquation := plotter.NewFunction(func(x float64) float64 {
		return theta0 + (theta1 * x)
	})
	linearEquation.Color = color.RGBA{B: 255, A: 255}

	s, err := plotter.NewScatter(coordsToXYs(dataset.data))
	handleError(err, "Error: could not add points of data")
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	xMax, yMax := getMaxValues(dataset.data)
	p.X.Min = 0
	p.X.Max = math.Floor(xMax * 1.10)
	p.Y.Min = 0
	p.Y.Max = math.Floor(yMax * 1.10)

	p.Add(s, linearEquation)
	err = p.Save(600, 600, "data_and_model_plotted.png")
	handleError(err, "Error: could not save plot of data and model")
}

func plotData(dataset Dataset, filename string) {
	p := plot.New()

	p.Title.Text = "Car price prediction"
	p.X.Label.Text = dataset.xName
	p.Y.Label.Text = dataset.yName

	s, err := plotter.NewScatter(coordsToXYs(dataset.data))
	handleError(err, "Error: could not add points of data")
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	p.Add(s)

	xMax, yMax := getMaxValues(dataset.data)
	p.X.Min = 0
	p.X.Max = math.Floor(xMax * 1.10)
	p.Y.Min = 0
	p.Y.Max = math.Floor(yMax * 1.10)

	err = p.Save(600, 600, filename)
	handleError(err, "Error: could not save plot of linear equation")
}
