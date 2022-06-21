package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// histPlot creates a histogram plot with the given number of samples.
func histPlot(samples int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var values plotter.Values
	var p *plot.Plot = plot.New()
	p.Title.Text = fmt.Sprintf("%d samples", samples)

	// Holds a normally distributed float64.
	var randValue float64

	// Creates a slice of random normally distributed values.
	for i := 0; i < samples; i++ {
		// Uses a mean of 50 and a standard deviation of 5.
		randValue = r.NormFloat64()*5 + 50
		values = append(values, randValue)
	}

	hist, err := plotter.NewHist(values, 30)
	if err != nil {
		panic(err)
	}
	p.Add(hist)

	fileName := "histogram_plot_" + strconv.Itoa(samples) + ".png"
	err = p.Save(5*vg.Inch, 5*vg.Inch, fileName)
	if err != nil {
		panic(err)
	}
}

// runSequential generates histogram plots without concurrency.
func runSequential() {
	repeat := 5
	var samples float64

	for i := 1; i <= repeat; i++ {
		// Starts with 10 samples and go up to 100,000.
		samples = math.Pow(10, float64(i))
		histPlot(int(samples))
	}
}

// runConcurrent generates histogram plots using goroutines.
func runConcurrent() {
	repeat := 5
	var samples float64

	// Creates a channel to receive confirmation from each goroutine.
	ch := make(chan int)

	// Creates a goroutine for each CPU.
	for i := 1; i <= repeat; i++ {
		samples = math.Pow(10, float64(i))
		go func(samples float64) {
			histPlot(int(samples))
			// Sends confirmation to the channel.
			ch <- 1
		}(samples)
	}

	// Receives confirmation from each goroutine.
	// This prevents the program from exiting before all goroutines have completed.
	for i := 1; i <= repeat; i++ {
		<-ch
	}
}

func main() {
	runConcurrent()
	// runSequential()
}
