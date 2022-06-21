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
	// Note that rand.NormFloat64() uses a shared global object with a Mutex lock on it.
	// Therefore, it shouldn't be used in the main goroutine.
	// Instead, allow each CPU to generate random values inside the goroutine.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < samples; i++ {
		// Generates random x and y values between 0.0 and 1.0.
		x, y := r.Float64(), r.Float64()
		// Uses the equation of a unit circle to check if the generated random sample fell inside the circle.
		if (x*x + y*y) < 1 {
			inside++
		}
	}
	return inside
}

// calcPiSequential calculates Pi without concurrency.
func calcPiSequential(samples int) float64 {
	inside := checkInsideCircle(samples)

	// Divides the count of samples inside the circle by total samples, which should equal to roughly π/4.
	ratio := float64(inside) / float64(samples)

	//  Multiplies by 4 to get our final estimate for π.
	return ratio * 4
}

// calcPiConcurrent calculates Pi using concurrent goroutines.
func calcPiConcurrent(samples int) float64 {
	// Determines the number of CPUs.
	var cpus = runtime.NumCPU()

	// Splits up the samples evenly among the CPUs.
	threadSamples := samples / cpus

	// Creates a channel to receive the results from each goroutine.
	results := make(chan float64)

	// Creates a goroutine for each CPU.
	for i := 0; i < cpus; i++ {
		go func() {
			inside := checkInsideCircle(threadSamples)
			// Sends result to the channel.
			results <- float64(inside) / float64(threadSamples) * 4
		}()
	}

	var total float64
	// Receives the results from each goroutine.
	for i := 0; i < cpus; i++ {
		total += <-results
	}

	return total / float64(cpus)
}

// displayOutput renders the result on the console.
func displayOutput(samples int, result float64) {
	fmt.Print("Samples: ", samples)
	// Adjusts spaces for formatting.
	spaces := 12 - int(math.Log10(float64(samples)))
	fmt.Printf("%s", strings.Repeat(" ", spaces))
	fmt.Println("Pi:", result)
}

func main() {
	// Starts with 10 samples and go up to 100,000,000.
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
