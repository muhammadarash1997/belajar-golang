package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Total when start", runtime.NumGoroutine())
	defer fmt.Println("Total when finish", runtime.NumGoroutine())
	errs := make(chan error)
	defer fmt.Println("main ended")
	
	go oneFunction(errs)
	go twoFunction(errs)
	go giverError(errs)
	fmt.Println("Total when run goroutine", runtime.NumGoroutine())
	<-errs
	time.Sleep(6 * time.Second)
}

func giverError(errs chan error) {
	time.Sleep(2 * time.Second)
	errs <- errors.New("error")
	errs <- errors.New("error")
	errs <- errors.New("error")
}

func oneFunction(errs chan error) {
	for {
		select {
		case <-errs:
			fmt.Println("oneFunction ended")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("oneFunction run")
		}
	}
}

func twoFunction(errs chan error) {
	for {
		select {
		case <-errs:
			fmt.Println("twoFunction ended")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("twoFunction run")
		}
	}
}
