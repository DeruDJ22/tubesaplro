package main

import "fmt"

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
			} else if sort == 2 {
				insertionSort(&name, n)
				dataTenant(name)
			} else {
				fmt.Println("Pilih angka yang benar")
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
			urutkanNama(&name, n)
			hapusTenant(&name, nama, &n)
		} else if a == 5 {
			var amount float64
			var jum_transaksi int
			fmt.Print("Masukkan nama tenant: ")
			fmt.Scan(&nama)
			fmt.Print("Masukkan jumlah transaksi: ")
			fmt.Scan(&jum_transaksi)
			for i := 0; i < jum_transaksi; i++ {
				fmt.Printf("Masukkan jumlah transaksi ke-%d: ", i+1)
				fmt.Scan(&amount)
				transaksi(&name, nama, n, amount)
			}
		} else if a == 6 {
			fmt.Println("Sampai jumpa lagi")
		} else {
			fmt.Println("Pilih angka yang benar")
		}
	}
}

// menambah data tenant
func tambahTenant(t *arrTenant, bnyk int, n *int) {
	startIndex := 0
	for startIndex < NMAX && (*t)[startIndex].nama != "" {
		startIndex++
	}
	if startIndex+bnyk > NMAX {
		fmt.Println("Tidak cukup ruang untuk menambah data baru")
	} else {
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

// menampilkan data tenant
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

// edit tenant
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

// hapus data tenant
func hapusTenant(name *arrTenant, tenant string, n *int) {
	index := binarySearch(*name, *n, tenant)
	if index == -1 {
		fmt.Println("Tenant tidak ditemukan")
	} else {
		fmt.Println("Data berhasil dihapus")
		for i := index; i < *n-1; i++ {
			name[i] = name[i+1]
		}
		name[*n-1] = Tenant{}
		*n--
	}
}

// menambah transaksi tenant
func transaksi(name *arrTenant, tenant string, n int, amount float64) {
	index := findTenant(*name, n, tenant)
	if index == -1 {
		fmt.Println("Tenant tidak ditemukan")
	} else {
		(*name)[index].transaksi++
		(*name)[index].pendapatan += amount
		adminRevenue := amount * 0.25
		tenantRevenue := amount * 0.75
		fmt.Printf("Transaksi dicatat untuk %s: Total %.2f (Admin: %.2f, Tenant: %.2f)\n", (*name)[index].nama, amount, adminRevenue, tenantRevenue)
		*&name[index].uAdmin += adminRevenue
		*&name[index].uTenant += tenantRevenue
	}
}

// mencari index tenant
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

// mencari index tenant dengan binary sort
func binarySearch(t arrTenant, n int, name string) int {
	kr := 0
	kn := n - 1
	for kr <= kn {
		med := (kr + kn) / 2
		if t[med].nama < name {
			kr = med + 1
		} else if t[med].nama > name {
			kn = med - 1
		} else {
			return med
		}
	}
	return -1
}

// mensorting nama tenant dengan selection sort
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

// mensorting nama tenant dengan insertion sort
func insertionSort(t *arrTenant, n int) {
	for i := 1; i < n; i++ {
		temp := t[i]
		j := i
		for j > 0 && temp.transaksi < t[j-1].transaksi {
			t[j] = t[j-1]
			j--
		}
		t[j] = temp
	}
}

// urut nama tenatn
func urutkanNama(t *arrTenant, n int) {
	for i := 1; i < n; i++ {
		temp := t[i]
		j := i
		for j > 0 && temp.nama < t[j-1].nama {
			t[j] = t[j-1]
			j--
		}
		t[j] = temp
	}
}

// cek nama tenant
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
