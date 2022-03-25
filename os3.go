// MENGEDIT ISI FILE

package main

import (
	"fmt"
	. "os"
)

var path = "D:/test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	// deteksi apakah file sudah ada
	var _, err = Stat(path)

	// buat file baru jika belum ada
	if IsNotExist(err) {
		var file, err = Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file berhasil dibuat", path)
}

func writeFile() {
	// buka file dengan level akses READ & WRITE
	var file, err = OpenFile(path, O_RDWR, 113) // O_RDWR adalah level akses dimana artinya Read dan Write, jadi kita bisa baca dan tulis/isi file tsb, 113 merupakan kode permission, setelah dicoba dengan berbagai macam angka tetap bisa, saya juga kurang tau mengapa
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("halo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("maribelajar golang\n")
	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> file berhasil di isi")
}

func main() {
	createFile()
	writeFile() // writeFile sama dijalankan berulang-ulang tidak akan merubah dan menambah, artinya isi file akan tetap seperti writeFile() pertama, sebenarnya dia mengoverwrite isi yg ada jadi tidak bisa menambah dengan writeString()
}

// DI BAWAH INI ADALAH LEVEL AKSES
// const (
//     // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
//     O_RDONLY int = syscall.O_RDONLY // open the file read-only.
//     O_WRONLY int = syscall.O_WRONLY // open the file write-only.
//     O_RDWR   int = syscall.O_RDWR   // open the file read-write.
//     // The remaining values may be or'ed in to control behavior.
//     O_APPEND int = syscall.O_APPEND // append data to the file when writing.
//     O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
//     O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
//     O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
//     O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
// )
