package main

import (
	"fmt"
	"sync"
	"testing"
)

var Wait sync.WaitGroup
var Mutex sync.Mutex
var Counter int = 0

func TestRaceConditions(t *testing.T) {
	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine)	// Race condition occured on here in which Counter variable get accessed by more than one goroutine without get locked first
	}

	Wait.Wait()

	fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {
	for count := 0; count < 2; count++ {
		value :=  Counter
		value++
		Mutex.Lock()
		Counter = value
		Mutex.Unlock()
	}

	Wait.Done()
}