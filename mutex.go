package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	counter int
	mutex   sync.Mutex
}

func (s *SafeCounter) Increment() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.counter++
}
func (s *SafeCounter) GetValue() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.counter
}

func main() {
	counter := SafeCounter{}

	numGoroutines := 3

	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				counter.Increment()
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}

	wg.Wait()
	finalValue := counter.GetValue()
	fmt.Printf("Final Counter Value: %d\n", finalValue)
}
