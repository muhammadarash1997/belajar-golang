// PROGRAM INI UNTUK MENCARI EIGEN VALUE ATAU EQUILIBRIUM ATAU NILAI KESETIMBANGAN
// DIMANA ADA DUA VARIABEL YAITU HUMAN DAN ZOMBIE YANG SELURUHNYA JIKA DIJUMLAHKAN MENJADI 300.
// SETIAP WAKTUNYA 20% HUMAN MENJADI ZOMBIE
// SETIAP WAKTUNYA 10% ZOMBIE MENJADI HUMAN
// DAN PROSES ZH (PERUBAHAN HUMAN KE ZOMBIE MAUPUN ZOMBIE KE HUMAN) AKAN TERUS TERJADI SAMPAI TITIK EQUILIBRIUM
// TITIK EQUILIBRIUM ADALAH KONDISI DIMANA HUMAN DAN ZOMBIE JUMLAHNYA TIDAK BERUBAH WALAUPUN TETAP ADA PROSES ZH
// TETAPI JUMLAH HUMAN TETAP SAMA DAN JUMLAH ZOMBIE TETAP SAMA

// MATERI INI ADA DI YT : https://www.youtube.com/watch?v=rowWM-MijXU

package main

import "fmt"

var A [][]float32 = [][]float32{
	{0.6, 0.3},
	{0.4, 0.7},
}

var ZH []float32 = []float32{
	20,
	280,
}

func main() {
	fmt.Println("Initial Value: ", ZH)
	before := ZH
	count := 0
	for {
		var xyZH []float32
		for i := 0; i < len(A); i++ {

			var xyTMP float32
			for j := 0; j < len(A[i]); j++ {
				xyTMP += A[i][j] * ZH[j]
			}
			xyZH = append(xyZH, xyTMP)
		}

		for i := 0; i < len(ZH); i++ {
			ZH[i] = xyZH[i]
		}

		fmt.Println("Result: ", ZH)

		for i := 0; i < len(before); i++ {
			if int(before[i]) != int(ZH[i]) {
				break
			}
			count++
		}
		if count > 100 {
			break
		}

		before = ZH
	}

	fmt.Println("Finish", ZH)
}
