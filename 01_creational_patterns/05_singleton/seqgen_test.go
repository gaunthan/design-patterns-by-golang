package singleton

import (
	"design-patterns-by-golang/01_creational_patterns/05_singleton/singleton"
	"sync"
)

func Example_one() {
	numRoutines := 1000000

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go func() {
			defer waitGroup.Done()

			singleton.GetInstance()
		}()
	}
	waitGroup.Wait()

	// Output:
	// creating instance
}
