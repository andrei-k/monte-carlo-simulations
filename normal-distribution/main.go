package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func histPlot(samples int) {
	// Note that rand.NormFloat64() uses a shared global object with a Mutex lock on it. Therefore, it shouldn't be used in the main goroutine. Instead, allow each CPU to generate random values inside the goroutine.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var values plotter.Values
	var p *plot.Plot = plot.New()
	p.Title.Text = fmt.Sprintf("%d samples", samples)

	// Generate a normally distributed float64
	var randValue float64

	// Create a slice of random normally distributed values
	for i := 0; i < samples; i++ {
		// Use a mean of 50 and a standard deviation of 5
		// NormFloat64()*desiredStdDev + desiredMean
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

	wg.Done()
}

var wg sync.WaitGroup

func main() {
	repeat := 5
	var samples float64

	for i := 1; i <= repeat; i++ {
		// Start with 10 samples and go up to 100,000
		samples = math.Pow(10, float64(i))
		go histPlot(int(samples))
		wg.Add(1)
	}

	wg.Wait()
}
