package main

import (
	"fmt"
)

const NMAX int = 10

type Tenant struct {
	nama            string
	transaksi       int
	pendapatan      float64
	uAdmin, uTenant float64
}

type arrTenant [NMAX]Tenant

func mulai() {
	var nTambah, a, n int
	var name arrTenant
	var nama string
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
			var sort int
			fmt.Println("___________________________")
			fmt.Println("Pilih kategori: ")
			fmt.Println("1. Penjualan terbanyak ke terdikit")
			fmt.Println("2. Penjualan terdikit ke terbanyak")
			fmt.Println("___________________________")
			fmt.Print("Pilih nomor: ")
			fmt.Scan(&sort)
			if sort == 1 {
				selectionSort(&name, n)
				dataTenant(name)
			} else {
				//else
			}
		} else if a == 3 {
			var namaBaru, namaLama string
			fmt.Print("Masukkan nama tenant lama: ")
			fmt.Scan(&namaLama)
			fmt.Print("Masukkan nama tenant baru: ")
			fmt.Scan(&namaBaru)
			editTenant(&name, namaBaru, namaLama, n)
		} else if a == 4 {
			fmt.Print("Masukkan nama tenant yang ingin dihapus: ")
			fmt.Scan(&nama)
			hapusTenant(&name, nama, &n)
		} else if a == 5 {
			//transaksi
		} else if a == 6 {
			fmt.Println("Sampai jumpa lagi")
		} else {
			fmt.Println("Pilih angka yang benar")
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
	} else {

		// Masukkan data ke dalam array
		for i := 0; i < bnyk; i++ {
			if startIndex+i < NMAX {
				var cek string
				fmt.Print("Nama Tenant: ")
				fmt.Scan(&cek)
				*n++
				if cekTenant(*t, cek) != -1 {
					fmt.Println("Data tenant sudah terdaftar")
					i--
				} else {
					fmt.Println("Nama tenant berhasil dibuat")
					(*t)[startIndex+i].nama = cek
				}
			}
		}
	}
}

func dataTenant(t arrTenant) {
	addDAta := false
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Println("Data tenant:")
	for i := 0; i < NMAX; i++ {
		if t[i].nama != "" {
			fmt.Printf("Tenant ke-%d: Nama Tenant = %s, Jumlah Transaksi = %d, Hasil Pendapatan = %.2f (Admin: %.2f, Tenant: %.2f)\n", i+1, t[i].nama, t[i].transaksi, t[i].pendapatan, t[i].uAdmin, t[i].uTenant)
			addDAta = true
		}
	}
	if addDAta == false {
		fmt.Println("Tidak ada data")
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
}

func editTenant(nama *arrTenant, namaBaru string, namaLama string, n int) {
	index := findTenant(*nama, n, namaLama)
	status := ""
	if index == -1 {
		status = "Tenant tidak ditemukan"
	} else if findTenant(*nama, n, namaBaru) != -1 {
		status = "Nama tenant sudah ada"
	} else if namaLama == namaBaru {
		status = "Nama tenant tidak berubah dikarenakan sama dengan nama tenant yang lama"
	} else {
		status = "Data berhasil di update!"
		nama[index].nama = namaBaru
	}
	fmt.Println(status)
}

func hapusTenant(nama *arrTenant, tenant string, n *int) {
	index := binarySort(*&nama, *n, tenant)
	status := ""
	if index == -1 {
		status = "Tenant tidak ditemukan"
	} else {
		status = "Data berhasil dihapus"
		for i := index; i < *n-1; i++ {
			nama[i] = nama[i+1]
		}
		*n--
	}
	fmt.Println(status)
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

func binarySort(T *arrTenant, n int, name string) int {
	var found int = -1
	var med int
	var kr int = 0
	var kn int = n - 1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if name > T[med].nama {
			kr = med + 1
		} else if name < T[med].nama {
			kn = med - 1
		} else {
			found = med
		}
	}
	return found
}

func selectionSort(T *arrTenant, n int) {
	for i := 0; i < n-1; i++ {
		idx_max := i
		for j := i + 1; j < n; j++ {
			if T[j].transaksi > T[idx_max].transaksi {
				idx_max = j
			}
		}
		T[i], T[idx_max] = T[idx_max], T[i]
	}
}

func insertionSort(T *arrTenant, n int) {
	var i, j int
	var temp Tenant
	i = 1
	for i <= n {
		temp = T[i]
		j = i
		for j > 0 && temp.transaksi < T[j-1].transaksi {
			T[j] = T[j-1]
			j--
		}
		T[j] = temp
	}
}

func cekTenant(T arrTenant, nama string) int {
	for i := 0; i < NMAX; i++ {
		if T[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	mulai()
}
