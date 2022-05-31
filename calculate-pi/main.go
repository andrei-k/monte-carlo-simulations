package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func calcPi(samples int) float64 {
	var inside int = 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < samples; i++ {
		// Generate a random x and y value between 0.0 and 1.0
		x, y := r.Float64(), r.Float64()
		// Use the equation of a unit circle to check if the generated random sample fell inside the circle
		if (x*x + y*y) < 1 {
			inside++
		}
	}

	// Divide the count of samples inside the circle by total samples, which should equal to roughly π/4
	ratio := float64(inside) / float64(samples)

	//  Multiply by 4 to get our final estimate for π
	return ratio * 4
}

// Calculate Pi using concurrent goroutines
func calcPiConcurrent(samples int, ch chan Channel) {
	// Send result to channel
	ch <- Channel{samples, calcPi(samples)}
}

// A struct for the channel to hold 2 values
type Channel struct {
	samples int
	result  float64
}

func main() {
	threads := 8
	var samples float64
	ch := make(chan Channel, threads)

	for i := 1; i <= threads; i++ {
		// Start with 10 samples and go up to 100,000,000
		samples = math.Pow(10, float64(i))
		go calcPiConcurrent(int(samples), ch)
	}

	fmt.Println("---------------------------------------------")
	fmt.Println("----------- Running concurrently ----------- ")
	fmt.Println("---------------------------------------------")

	for i := 0; i < threads; i++ {
		output := <-ch // Receive result from channel
		// TODO: render output as a function

		fmt.Print("Samples: ", output.samples)
		// Adjust spaces for formatting
		fmt.Printf("%s", strings.Repeat(" ", 12-int(math.Log10(float64(output.samples)))))
		fmt.Println("Pi:", output.result)
	}

	fmt.Println("---------------------------------------------")
	fmt.Println("----------- Running sequentially ----------- ")
	fmt.Println("---------------------------------------------")

	for i := 1; i <= threads; i++ {
		// Start with 10 samples and go up to 100,000,000
		samples = math.Pow(10, float64(i))
		// TODO: render output as a function

		fmt.Print("Samples: ", int(samples))
		// Adjust spaces for formatting
		fmt.Printf("%s", strings.Repeat(" ", 12-i))
		fmt.Println("Pi:", calcPi(int(samples)))
	}
}
