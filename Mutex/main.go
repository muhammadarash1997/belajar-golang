package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	Name    string
	lock    sync.Mutex
}

func (a *Account) Withdraw(amount int, wg *sync.WaitGroup) {
    defer wg.Done()
    a.lock.Lock()
	time.Sleep(2 * time.Second)
    a.balance -= amount
    a.lock.Unlock()
}

func (a *Account) Deposit(amount int, wg *sync.WaitGroup) {
    defer wg.Done()
    a.lock.Lock()
	time.Sleep(2 * time.Second)
    a.balance += amount
    a.lock.Unlock()
}

func (a *Account) GetBalance(wg *sync.WaitGroup) int {
	defer wg.Done()
    // a.lock.Lock()
    // defer a.lock.Unlock()

    return a.balance
}

func main() {
    var account Account
    var wg sync.WaitGroup

    account.Name = "Test account"

    for i := 0; i < 20; i++ {
        wg.Add(1)
        go account.Deposit(100, &wg)
    }

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go account.Withdraw(100, &wg)
    }

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go fmt.Printf("Balance: %d\n", account.GetBalance(&wg))
	}
	
    wg.Wait()
	fmt.Println("Finished")
}
