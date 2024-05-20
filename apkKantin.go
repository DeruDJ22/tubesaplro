package main

import (
	"fmt"
)

const NMAX int = 10

type Tenant struct {
	nama       string
	transaksi  int
	pendapatan float64
}

type arrTenant [NMAX]Tenant

func mulai() {
	var nTambah, a, n int
	var name arrTenant
	for a != 6 {
		fmt.Println("________________________")
		fmt.Println("Selamat Datang Kembali")
		fmt.Println("1. Tambah Tenant")
		fmt.Println("2. Data tenant")
		fmt.Println("3. Edit Tenant")
		fmt.Println("4. Hapus Tenant")
		fmt.Println("5. Catat Transaksi")
		fmt.Println("6. Keluar")
		fmt.Println("________________________")
		fmt.Print("Pilih nomor: ")
		fmt.Scan(&a)
		if a == 1 {
			fmt.Print("Berapa banyak data yang ingin ditambah: ")
			fmt.Scan(&nTambah)
			tambahTenant(&name, nTambah, &n)
		} else if a == 2 {
			dataTenant(name)
		} else if a == 3 {
			var namaBaru, namaLama string
			fmt.Print("Masukkan nama tenant lama: ")
			fmt.Scan(&namaLama)
			fmt.Print("Masukkan nama tenant baru: ")
			fmt.Scan(&namaBaru)
			editTenant(name, namaBaru, namaLama, n)
		}
	}
}

// Fungsi untuk menambah tenant
func tambahTenant(t *arrTenant, bnyk int, n *int) {
	startIndex := 0
	for startIndex < NMAX && (*t)[startIndex].nama != "" {
		startIndex++
	}

	// Jika array tidak cukup besar untuk menampung data baru, beri pesan kesalahan
	if startIndex+bnyk > NMAX {
		fmt.Println("Tidak cukup ruang untuk menambah data baru")
		return
	}

	// Masukkan data ke dalam array
	for i := 0; i < bnyk; i++ {
		if startIndex+i < NMAX {
			fmt.Print("Nama Tenant: ")
			fmt.Scan(&t[startIndex+i].nama)
		}
		*n++
	}
}

func dataTenant(t arrTenant) {
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println("Data tenant:")
	for i := 0; i < NMAX; i++ {
		if t[i].nama != "" {
			fmt.Printf("Tenant ke-%d: Nama Tenant = %s, Jumlah Transaksi = %d, Hasil Pendapatan = %.2f\n", i+1, t[i].nama, t[i].transaksi, t[i].pendapatan)
		}
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
}

func editTenant(nama arrTenant, namaBaru, namaLama string, n int) {
	index := findTenant(nama, n, namaLama)
	if index == -1 {
		fmt.Println("Tenant tidak ditemukan")
		return
	}
	if findTenant(nama, n, namaBaru) != -1 {
		fmt.Println("Nama tenant sudah ada")
		return
	}
	fmt.Println("Data berhasil di update!")
	nama[index].nama = namaBaru
}

func findTenant(T arrTenant, n int, name string) int {
	var found int = -1
	var i int = 0
	for i < n && found == -1 {
		if T[i].nama == name {
			found = i
		}
		i++
	}
	return found
}

func main() {
	mulai()
}
