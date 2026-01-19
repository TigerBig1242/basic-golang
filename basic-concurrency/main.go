package main

import (
	"fmt"
	"sync"
	"time"
)

// type worker struct {
// 	id int
// }

// type job struct {
// 	id int
// }

func RunWorkers(numWorkers int, numJobs int) {
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	for work := 1; work <= numWorkers; work++ {
		wg.Add(1)
		go worker(work, jobs, &wg)
		// go worker(work, jobs, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)
	wg.Wait()
	fmt.Printf("All job finnished")
}

func worker(workerID int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for jobID := range jobs {
		fmt.Printf("Worker ID:[%d] processing job ID:[%d]\n", workerID, jobID)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	RunWorkers(2, 10)
	// Create channel send jobs wo worker
	// jobs := make(chan int, 5)

	// // wg it's like a check point fir each loop
	// var wg sync.WaitGroup

	// wg.Add(1)
	// worker(1, jobs, &wg)
	// // go worker(2, jobs, &wg)
	// // go worker(3, jobs, &wg)

	// fmt.Printf("Main: sending test jobs...")
	// jobs <- 101
	// // jobs <- 102
	// // jobs <- 103
	// // jobs <- 104
	// // jobs <- 105
	// // jobs <- 106
	// // jobs <- 107
	// // jobs <- 108
	// // jobs <- 109
	// // jobs <- 110

	// close(jobs)

	// wg.Wait()

	// fmt.Printf("Main: completed successfully")

	// var wg sync.WaitGroup

	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Printf("routine 1: %d\n", i)
	// 	}
	// }()
	// // time.Sleep(1 * time.Millisecond)

	// go func() {
	// 	defer wg.Done()
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Printf("routine 2: %d\n", i)
	// 	}
	// }()
	// time.Sleep(2 * time.Millisecond)

	// wg.Wait()
}
