package main

import (
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func HistPlot(values plotter.Values) {
	p := plot.New()
	p.Title.Text = "Histogram Plot"

	hist, err := plotter.NewHist(values, 30)
	if err != nil {
		panic(err)
	}
	p.Add(hist)

	err = p.Save(5*vg.Inch, 5*vg.Inch, "histogram.png")
	if err != nil {
		panic(err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var values plotter.Values

	// Generate a normally distributed float64
	// sample = NormFloat64()*desiredStdDev + desiredMean
	var sample float64

	// Create a slice of 100,000 random normally distributed samples
	for i := 0; i < 100000; i++ {
		// Use a mean of 50 and a standard deviation of 5
		sample = rand.NormFloat64()*5 + 50
		values = append(values, sample)
	}

	HistPlot(values)
}
