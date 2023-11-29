package main
import (
	"fmt"
	"sync"
	"time"
)
func libraryReader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Reader %d has started reading a book\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Reader %d has finished reading a book\n", id)
}
func main() {
	const readersNumber = 5
	var wg sync.WaitGroup
	wg.Add(readersNumber)
	for i := 1; i <= readersNumber; i++ {
		go libraryReader(i, &wg)
	}
	wg.Wait()
}
