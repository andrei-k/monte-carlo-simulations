# Calculate the value of π using a Monte Carlo Simulation

This example was taken from this great [article](https://www.soroushjp.com/2015/02/07/go-concurrency-is-not-parallelism-real-world-lessons-with-monte-carlo-simulations/) that taught me a lot. I rewrote the code in my own way though.

The program uses separate goroutines executing each set of samples. A channel (type struct with 2 data fields) is used to communicate the results of each goroutine to the main thread.

The `GOMAXPROCS` no longer needs to be explicitly set to use all available CPUs, as mentioned in the article. As of Go 1.5, programs run with `GOMAXPROCS` set to the number of cores available by default, whereas in prior releases it defaulted to 1.

## Usage

Run in Terminal

```go
go run .
```
