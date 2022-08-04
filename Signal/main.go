package main

import (
	"fmt"
	"os"
	"os/signal"
)

// Golang's os/signal the package allows you to configure the behavior of your Golang program upon receiving certain types of UNIX or LINUX signals.

func main() {
	sigs := make(chan os.Signal)

	// Passing channel ke Notify yang mana di dalam Notify ada goroutine
	signal.Notify(sigs)

	done := make(chan bool)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("existing")
}

// OUTPUT
// awaiting signal
// ^C atau ctrl+c
// interrupt
// existing

// Console akan menampilkan awaiting signal di console dan terus menunggu sampai kita melakukan exit di terminal (ctrl+c)
// Setelah exit di terminal maka console akan menampilkan interrupt dan existing

// REAL CASE
// Real case nyata misal mau cleanup resource (tutup koneksi db, ngeflush file handle,
// atau nunggu background job kelar dll) pas aplikasi shutdown. Ini istilahnya graceful shutdown.
