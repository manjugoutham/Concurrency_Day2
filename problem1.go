/* Write two goroutines which have a race condition when executed concurrently */

package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incrementCounter(1)
	go incrementCounter(2)

	wg.Wait()

	fmt.Printf("Final counter: %d\n", counter)
}

func incrementCounter(id int) {
	for i := 0; i < 5; i++ {
		value := counter
		value++
		fmt.Printf("Goroutine id %d: Counter value: %d\n", id, value)
		counter = value
	}
	wg.Done()
}
