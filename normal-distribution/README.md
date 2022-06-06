# Normal Distribution Simulation

Go program running the example outlined in this [article](https://machinelearningmastery.com/monte-carlo-sampling-for-probability/).  

We will use a Gaussian distribution with a mean of 50 and a standard deviation of 5 and draw random samples from this distribution.

Let’s pretend we don’t know the form of the probability distribution for this random variable and we want to sample the function to get an idea of the probability density. We can draw a sample of a given size and plot a histogram to estimate the density.

To make the example more interesting, we will repeat this experiment four times with different sized samples. We would expect that as the size of the sample is increased, the probability density will better approximate the true density of the target function, given the law of large numbers.

We can see that the small sample sizes of 10 and 50 do not effectively capture the density of the target function. We can see that 100 samples is better, but it is not until 1,000 samples that we clearly see the familiar bell-shape of the Gaussian probability distribution.

This highlights the need to draw many samples, even for a simple random variable, and the benefit of increased accuracy of the approximation with the number of samples drawn.

## Usage

Run in Terminal

```go
go run .
```
