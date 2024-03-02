package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func getDataTeman(absen int) *Teman {
	dataTeman := map[int]Teman{
		1: {"John Doe", "Jakarta", "Software Engineer", "Go diketahui memiliki kinerja yang sangat baik, terutama dalam konteks konkurensi dan pemrograman paralel"},
		2: {"Jane Smith", "Bandung", "Data Scientist", "Go didesain mudah dipahami dan digunakan, dengan sintaks yang sederhana namun kuat"},
		3: {"Michael Johnson", "Surabaya", "Web Developer", "Meskipun relatif baru, komunitas Go berkembang pesat"},
		4: {"Diego", "Riau", "Data Analyst", "Go sering digunakan dalam pengembangan aplikasi cloud dan infrastruktur, terutama karena kemampuannya dalam menangani beban kerja tinggi dan skala yang besar"},
		5: {"Marchel", "Bali", "Backhand Developer", "Go didesain dengan keamanan dan pemeliharaan dalam pikiran"},
	}

	teman, found := dataTeman[absen]
	if !found {
		return nil
	}
	return &teman
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Gunakan: go run biodata.go <nomor_absen>")
		return
	}
	absen := os.Args[1]
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	teman := getDataTeman(absenInt)
	if teman == nil {
		fmt.Println("Tidak terdapat nama teman dengan nomer absen tersebut")
		return
	}

	fmt.Println("Data Teman:")
	fmt.Println("Nama :", teman.Nama)
	fmt.Println("Alamat :", teman.Alamat)
	fmt.Println("Pekerjaan :", teman.Pekerjaan)
	fmt.Println("Alasan :", teman.Alasan)
}
