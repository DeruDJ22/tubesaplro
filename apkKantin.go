package main

import (
	"fmt"
)

const NMAX int = 10

type Tenant struct {
	Name         string
	Transactions int
	Revenue      float64
}

type arrTenant [NMAX]Tenant

func main() {
	var choice int
	var name arrTenant
	for {
		fmt.Println("------------------------------")
		fmt.Println("Pilih operasi:")
		fmt.Println("1. Tambah Tenant")
		fmt.Println("2. Data tenant")
		fmt.Println("3. Edit Tenant")
		fmt.Println("4. Hapus Tenant")
		fmt.Println("5. Catat Transaksi")
		fmt.Println("6. Keluar")
		fmt.Println("------------------------------")
		fmt.Print("Masukkan pilihan (1-6): ")

		// Membaca pilihan pengguna
		fmt.Scan(&choice)

		if choice == 1 {
			var name arrTenant
			var banyak int
			fmt.Print("Berapa banyak tenant yang ingin di buat? ")
			fmt.Scan(&banyak)
			addTenant(&name, banyak)
		} else if choice == 2 {
			displayTenants(name)
		} else if choice == 3 {
			// edit
		} else if choice == 4 {
			// hapus
		} else if choice == 5 {
			// catat transaksi
		} else if choice == 6 {
			// keluar
		}
	}
}

// Fungsi untuk menambah tenant
func addTenant(name *arrTenant, bnyk int) {
	startIndex := 0
	for startIndex < NMAX && (*name)[startIndex].Name != "" {
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
			fmt.Scan(&name[startIndex+i].Name)
		}
	}
}
func displayTenants(A arrTenant) {
	fmt.Println()
	fmt.Println("Data makanan:")
	for i, m := range A {
		if m.Name != "" {
			fmt.Printf("| Makanan ke-%d: Nama = %s, Kategori = %s, Harga = %d | \n", i+1, m.Name, m.Revenue, m.Transactions)
		}
	}
}
