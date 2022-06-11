package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

func checkInsideCircle(samples int) (inside int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < samples; i++ {
		// Generate a random x and y value between 0.0 and 1.0
		x, y := r.Float64(), r.Float64()
		// Use the equation of a unit circle to check if the generated random sample fell inside the circle
		if (x*x + y*y) < 1 {
			inside++
		}
	}
	return inside
}

func calcPiSequential(samples int) float64 {
	inside := checkInsideCircle(samples)

	// Divide the count of samples inside the circle by total samples, which should equal to roughly π/4
	ratio := float64(inside) / float64(samples)

	//  Multiply by 4 to get our final estimate for π
	return ratio * 4
}

// Calculate Pi using concurrent goroutines
func calcPiConcurrent(samples int) float64 {
	// Determine the number of CPUs
	var cpus = runtime.NumCPU()

	// Split up the samples evenly among the CPUs
	threadSamples := samples / cpus

	// Create a channel to receive the results from each goroutine
	results := make(chan float64)

	// Create a goroutine for each CPU
	for i := 0; i < cpus; i++ {
		go func() {
			inside := checkInsideCircle(threadSamples)
			results <- float64(inside) / float64(threadSamples) * 4
		}()
	}

	var total float64
	for i := 0; i < cpus; i++ {
		total += <-results
	}

	return total / float64(cpus)
}

// Render result to console
func displayOutput(samples int, result float64) {
	fmt.Print("Samples: ", samples)
	// Adjust spaces for formatting
	spaces := 12 - int(math.Log10(float64(samples)))
	fmt.Printf("%s", strings.Repeat(" ", spaces))
	fmt.Println("Pi:", result)
}

func main() {
	// Start with 10 samples and go up to 100,000,000
	iterations := 8
	var samples float64

	fmt.Println("---------------------------------------------")
	fmt.Println("--------- Calculating sequentially --------- ")
	fmt.Println("---------------------------------------------")

	for i := 1; i <= iterations; i++ {
		samples = math.Pow(10, float64(i))
		displayOutput(int(samples), calcPiSequential(int(samples)))
	}

	fmt.Println("---------------------------------------------")
	fmt.Println("--------- Calculating concurrently --------- ")
	fmt.Println("---------------------------------------------")

	for i := 1; i <= iterations; i++ {
		samples = math.Pow(10, float64(i))
		displayOutput(int(samples), calcPiConcurrent(int(samples)))
	}
}
