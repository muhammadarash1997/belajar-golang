package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	Name    string
	lock    sync.RWMutex
}

func (a *Account) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.lock.Lock()
	time.Sleep(2 * time.Second)
	defer a.lock.Unlock()
	a.balance -= amount

}

func (a *Account) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	a.lock.Lock()
    time.Sleep(2 * time.Second)
	defer a.lock.Unlock()
	a.balance += amount
}

func (a *Account) GetBalance(wg *sync.WaitGroup) int {
	defer wg.Done()
	a.lock.RLock()
	defer a.lock.RUnlock()

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

// NOTE : UNTUK MELIHAT PERBEDAAN ANTARA MUTEX DAN RWMUTEX
// MAKA DAPAT DIGANTI PADA LINE 36 & 37 MEJADI
// DENGAN LOCK & UNLOCK SAJA JIKA UNTUK MUTEXT
// DENGAN RLOCK & RUNLOCK SAJA JIKA UNTUK RWMUTEXT

// KEKURANGAN DENGAN MUTEX ADALAH ANTARA READ & WRITE SALING ANTRI
// SEMENTARA DENGAN RWMUTEX ANTARA READ & WRITE MEMILIKI ANTRIAN TERPISAH