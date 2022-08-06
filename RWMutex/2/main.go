package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int = 0
var wg sync.WaitGroup
var lock sync.RWMutex

func write() {
	defer wg.Done()

	lock.Lock()
	defer lock.Unlock()

	time.Sleep(3 * time.Second)

	balance += 10

}

func read() {
	defer wg.Done()

	time.Sleep(1 * time.Second)

	lock.RLock()
	defer lock.RUnlock()
	
	fmt.Println(balance)
}

func main() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go write()
	}

	wg.Add(1)
	go read()

	wg.Wait()
	fmt.Println("Finished")
}
