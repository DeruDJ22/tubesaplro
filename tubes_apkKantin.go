package main

import "fmt"

const NMAX int = 10

type makanan struct {
	nama, kategori string
	harga          int
}
type arrMkn [NMAX]makanan

func mulai() {
	var a, nTambah int
	var tambah arrMkn
	for a != 5 {
		fmt.Println("________________________")
		fmt.Println("Selamat Datang Kembali")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Lihat Data")
		fmt.Println("3. Edit Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Keluar")
		fmt.Println("________________________")
		fmt.Print("Pilih nomor: ")
		fmt.Scan(&a)
		if a == 1 {
			fmt.Println("Berapa banyak data yang ingin ditambah?")
			fmt.Scan(&nTambah)
			bacaData(&tambah, nTambah)
		} else if a == 2 {
			cetakData(tambah)
		}
	}
}

func bacaData(A *arrMkn, B int) {
	// Temukan indeks pertama yang kosong tanpa menggunakan break
	startIndex := 0
	for startIndex < NMAX && (*A)[startIndex].nama != "" {
		startIndex++
	}

	// Jika array tidak cukup besar untuk menampung data baru, beri pesan kesalahan
	if startIndex+B > NMAX {
		fmt.Println("Tidak cukup ruang untuk menambah data baru")
		return
	}

	// Masukkan data ke dalam array
	for i := 0; i < B; i++ {
		if startIndex+i < NMAX {
			fmt.Print("Nama makanan, kategori dan harga makanan: ")
			fmt.Scan(&(*A)[startIndex+i].nama, &(*A)[startIndex+i].kategori, &(*A)[startIndex+i].harga)
		}
	}
}

func cetakData(A arrMkn) {
	fmt.Println()
	fmt.Println("Data makanan:")
	for i, m := range A {
		if m.nama != "" {
			fmt.Printf("| Makanan ke-%d: Nama = %s, Kategori = %s, Harga = %d | \n", i+1, m.nama, m.kategori, m.harga)
		}
	}
}

func main() {
	//mulai
	mulai()
}
