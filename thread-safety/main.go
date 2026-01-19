package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (s *SafeCounter) Inc(workerID int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(workerID) * 100 * time.Millisecond)
	s.mu.Lock()
	s.count++
	defer s.mu.Unlock()
	// defer wg.Done()
	// fmt.Println("workerID :", id)
}

func (s *SafeCounter) Value() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup
	totalWorkers := 100
	fmt.Println("เริ่มให้คนงานทำงาน...")
	for i := 0; i < totalWorkers; i++ {
		wg.Add(1)
		go counter.Inc(i, &wg)
	}

	stopMonitoring := make(chan bool)
	go func() {
		lastSeenValue := -1
		for {
			select {
			case <-stopMonitoring:
				return
			default:
				currentValue := counter.Value()
				if currentValue > lastSeenValue {
					fmt.Printf("ขณะนี้ค่าปัจจุบันคือ :%d\n", counter.Value())
					lastSeenValue = currentValue
				}

				if currentValue == totalWorkers {
					return
				}
				time.Sleep(20 * time.Millisecond)
				// time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			}
		}
	}()
	wg.Wait()
	stopMonitoring <- true
	fmt.Println("last current value :", counter.Value())
}
