package main

import "fmt"

type Latihan struct {
	Nama    string
	Jenis   string
	Durasi  int
	Kalori  int
	Tanggal string
}

var daftarLatihan []Latihan
var pengguna, kataSandi string

func login() {
	fmt.Print("Masukkan Username: ")
	fmt.Scanln(&pengguna)

	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&kataSandi)

	fmt.Printf("\nLogin berhasil! Selamat datang, %s!\n\n", pengguna)
}

func tambahLatihan() {
	fmt.Println("\n--- Tambah Latihan ---")
	var l Latihan

	fmt.Print("Nama Latihan: ")
	fmt.Scanln(&l.Nama)

	fmt.Print("Jenis Latihan: ")
	fmt.Scanln(&l.Jenis)

	fmt.Print("Durasi (menit): ")
	fmt.Scanln(&l.Durasi)

	fmt.Print("Kalori terbakar: ")
	fmt.Scanln(&l.Kalori)

	fmt.Print("Tanggal (YYYY-MM-DD): ")
	fmt.Scanln(&l.Tanggal)

	daftarLatihan = append(daftarLatihan, l)
	fmt.Println("Latihan berhasil ditambahkan!\n")
}

func ubahLatihan() {
	fmt.Println("\n--- Ubah Latihan ---")
	var nama string
	fmt.Print("Masukkan nama latihan yang ingin diubah: ")
	fmt.Scanln(&nama)

	for i, l := range daftarLatihan {
		if l.Nama == nama {
			fmt.Println("Data ditemukan. Masukkan data baru.")

			fmt.Print("Nama Latihan: ")
			fmt.Scanln(&daftarLatihan[i].Nama)

			fmt.Print("Jenis Latihan: ")
			fmt.Scanln(&daftarLatihan[i].Jenis)

			fmt.Print("Durasi (menit): ")
			fmt.Scanln(&daftarLatihan[i].Durasi)

			fmt.Print("Kalori terbakar: ")
			fmt.Scanln(&daftarLatihan[i].Kalori)

			fmt.Print("Tanggal (YYYY-MM-DD): ")
			fmt.Scanln(&daftarLatihan[i].Tanggal)

			fmt.Println("Latihan berhasil diubah!\n")

			return
		}
	}
	fmt.Println("Latihan tidak ditemukan.")
}

func hapusLatihan() {
	fmt.Println("\n--- Hapus Latihan ---")
	var nama string
	fmt.Print("Masukkan nama latihan yang ingin dihapus: ")
	fmt.Scanln(&nama)

	for i, l := range daftarLatihan {
		if l.Nama == nama {
			daftarLatihan = append(daftarLatihan[:i], daftarLatihan[i+1:]...)
			fmt.Println("Latihan berhasil dihapus.")
			return
		}
	}
	fmt.Println("Latihan tidak ditemukan.")
}

func cariLatihanSequential() {
	fmt.Println("\n--- Cari Latihan (Sequential Search) ---")
	var jenis string
	fmt.Print("Masukkan jenis latihan: ")
	fmt.Scanln(&jenis)

	ditemukan := false
	for _, l := range daftarLatihan {
		if l.Jenis == jenis {
			fmt.Printf("%+v\n", l)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Latihan tidak ditemukan.")
	}
}

func urutkanDurasiSelection() {
	n := len(daftarLatihan)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if daftarLatihan[j].Durasi < daftarLatihan[min].Durasi {
				min = j
			}
		}
		daftarLatihan[i], daftarLatihan[min] = daftarLatihan[min], daftarLatihan[i]
	}
	fmt.Println("Data latihan diurutkan berdasarkan durasi:")
	for _, l := range daftarLatihan {
		fmt.Printf("%+v\n", l)
	}
}

func urutkanKaloriInsertion() {
	for i := 1; i < len(daftarLatihan); i++ {
		kunci := daftarLatihan[i]
		j := i - 1
		for j >= 0 && daftarLatihan[j].Kalori > kunci.Kalori {
			daftarLatihan[j+1] = daftarLatihan[j]
			j--
		}
		daftarLatihan[j+1] = kunci
	}
	fmt.Println("Data latihan diurutkan berdasarkan kalori:")
	for _, l := range daftarLatihan {
		fmt.Printf("%+v\n", l)
	}
}

func tampilkanLaporan() {
	fmt.Println("\n--- Laporan Latihan ---")
	n := len(daftarLatihan)
	awal := 0
	if n > 10 {
		awal = n - 10
	}
	fmt.Println("10 Aktivitas Terakhir:")
	for i := awal; i < n; i++ {
		fmt.Printf("%+v\n", daftarLatihan[i])
	}

	totalKalori := 0
	for _, l := range daftarLatihan {
		totalKalori += l.Kalori
	}
	fmt.Printf("Total Kalori Terbakar: %d\n", totalKalori)
}

func rekomendasiLatihan() {
	fmt.Println("\n--- Rekomendasi Latihan ---")
	jumlah := make(map[string]int)
	for _, l := range daftarLatihan {
		jumlah[l.Jenis]++
	}

	var jenisTerbanyak string
	terbanyak := 0
	for jenis, jumlahLatihan := range jumlah {
		if jumlahLatihan > terbanyak {
			terbanyak = jumlahLatihan
			jenisTerbanyak = jenis
		}
	}

	if jenisTerbanyak != "" {
		fmt.Printf("Rekomendasi: Lakukan lebih banyak '%s' karena sering dilakukan.\n", jenisTerbanyak)
	} else {
		fmt.Println("Belum ada data latihan untuk direkomendasikan.")
	}
}

func main() {
	login()

	for {
		fmt.Println(`
====== MENU ======
1. Tambah Latihan
2. Ubah Latihan
3. Hapus Latihan
4. Cari Latihan (Sequential)
5. Urutkan Latihan (Durasi - Selection Sort)
6. Urutkan Latihan (Kalori - Insertion Sort)
7. Laporan Latihan
8. Rekomendasi Latihan
0. Keluar
==================`)
		var pilihan string
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			tambahLatihan()
		case "2":
			ubahLatihan()
		case "3":
			hapusLatihan()
		case "4":
			cariLatihanSequential()
		case "5":
			urutkanDurasiSelection()
		case "6":
			urutkanKaloriInsertion()
		case "7":
			tampilkanLaporan()
		case "8":
			rekomendasiLatihan()
		case "0":
			fmt.Println("Terima kasih telah menggunakan aplikasi kami, Sampai Jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
