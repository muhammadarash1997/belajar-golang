// NOTE : JIKA INGIN PELAJARI GOROUTINE LENGKAP ADA DI "DASAR PEMROGRAMAN GOLANG" PDF

// Unbuffered Channel dan Buffered Channel sebenarnya sama saja
// Mereka bisa diibaratkan int dan []int
// Unbuffered hanya bisa menampung 1 data sedangkan Buffered bisa 1 hingga lebih terantung set di awal

package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(2)

	// 2 adalah jumlah yang dapat ditampung channel ini
	// 2 artinya (0, 1, 2) berarti ada 3 data yang dapat ditampung
	// Jika kita set 0 maka hanya 1 saja yang dapat ditampung yang mana artinya menjadi unbuffered channel karena hanya 1
	messages := make(chan int, 2)

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}

	fmt.Scanln()
	// Agar fungsi main() tidak langsung berhenti, karena kita ingin menunggu agar semua fungsi goroutine selesai semua.
	// Karena jika fungsi main() berhenti maka semua fungsi apapun akan berhenti bahkan goroutine sekalipun.
	// Fungsi main() tidak akan berhenti sampai kita menekan tombol enter di console.
}

// Pada func() yang bertugas mencetak receive data kita tidak menggunakan parameter
// channel pada fungsi tsb, karena variabel channel message berada dalam scope yang sama.
// Apabila func() tsb berada di luar func main() atau berada di scope yang berbeda dengan
// variabel message, maka pada func() harus menerima parameter bertipe channel.