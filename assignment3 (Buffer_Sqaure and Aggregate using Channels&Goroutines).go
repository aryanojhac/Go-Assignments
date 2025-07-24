// Write a code to squares numbers from a list concurrently and aggregates the squared values.
package main

import (
	"fmt"
	"sync"
)

func squareWorker(number int, waitgroup *sync.WaitGroup, resultChan chan int) {
	defer waitgroup.Done() // Tell WaitGroup when done with this goroutine (Decreases the counter)
	square := number * number
	resultChan <- square
}

func aggregateSquare(ch chan int, done chan bool) {
	totalSum := 0
	for square := range ch {
		totalSum += square
	}
	fmt.Println("Total of the Squares is: ", totalSum)
	done <- true
}

func main() {
	var numberByUser uint
	fmt.Println("Enter the number of Digits: ")
	_, err := fmt.Scanln(&numberByUser)
	if err != nil {
		fmt.Println("Invalid Number", err)
		return
	}
	if numberByUser == 0 {
		fmt.Println("Number of Digits must be greater than 0.")
		return
	}

	fmt.Println("Enter the numbers of which you want agregate of squares:")
	numbers := make([]int, numberByUser)
	for i := range numbers {
		_, err := fmt.Scanln(&numbers[i])
		if err != nil {
			fmt.Println("Cannot take the number as input")
			return
		}
	}

	var wg sync.WaitGroup
	resultChan := make(chan int, len(numbers)) // Channel to hold the results
	done := make(chan bool)
	go aggregateSquare(resultChan, done) // Starting the Goroutine

	// Iterating over the number and starting goroutine for each
	for _, num := range numbers {
		wg.Add(1)                             // Tell Waitgroup to wait for another goroutine
		go squareWorker(num, &wg, resultChan) // starting the goroutine
	}

	wg.Wait() //wait until the counter becomes 0
	close(resultChan)

	<-done
}
