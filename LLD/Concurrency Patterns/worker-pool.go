package main

/*
The worker pool is a really powerful pattern that lets us distributes the work across
multiple workers (goroutines) concurrently.

In our example, we have a jobs channel to which we will send our jobs
and a results channel where our workers will send the results once they've finished doing the work.

After that, we can launch our workers concurrently and simply receive the results from the results channel.

Ideally, totalWorkers should be set to runtime.NumCPU() which gives us the number of
logical CPUs usable by the current process.
*/
import (
	"fmt"
)

const totalJobs = 10
const totalWorkers = 4

func main() {
	jobs := make(chan int, totalJobs)
	results := make(chan int)

	// 1. Start workers
	for w := 1; w <= totalWorkers; w++ {
		go worker(w, jobs, results)
	}

	// 2. Send jobs
	for j := 1; j <= totalJobs; j++ { // This loop is non blocking as jobs is buffered channel
		jobs <- j
	}
	close(jobs) // Closing tells workers no more jobs are coming

	// 3. Collect results
	// This loop blocks until all results are received
	for a := 1; a <= totalJobs; a++ {
		res := <-results
		fmt.Printf("Result received: %d\n", res)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	// Each worker pulls a job, processes it, then loops back for the next.
	// No new goroutines are spawned here.
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)

		// Simulate work
		result := j * 2
		results <- result

		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}
