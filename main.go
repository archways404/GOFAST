package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a random array of numbers
	array := generateRandomArray(1000000, 1, 1000000)

	// Measure the start time
	startTime := time.Now()

	// Find a random number within the array
	randomNumber := findRandomNumber(array)

	// Measure the end time
	endTime := time.Now()

	// Calculate the execution time
	executionTime := endTime.Sub(startTime)

	fmt.Printf("Random number within the array: %d\n", randomNumber)
	fmt.Printf("Execution time: %v\n", executionTime)
}

// generateRandomArray generates a random array of numbers with the specified length and range
func generateRandomArray(length, min, max int) []int {
	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = rand.Intn(max-min+1) + min
	}
	return array
}

// findRandomNumber finds a random number within the array
func findRandomNumber(array []int) int {
	index := rand.Intn(len(array))
	return array[index]
}
