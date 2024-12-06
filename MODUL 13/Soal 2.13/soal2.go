package main

import "fmt"

const nMax = 7919

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	tahun     int
	rating    int
	eksemplar int
}

type DaftarBuku [nMax]Buku

var pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	for i := 0; i < *n; i++ {
		var id, judul, penulis, penerbit string
		var tahun, rating, eksemplar int
		fmt.Print("Masukkan data buku (id judul penulis penerbit tahun rating eksemplar): ")
		fmt.Scan(&id, &judul, &penulis, &penerbit, &tahun, &rating, &eksemplar)
		(*pustaka)[i] = Buku{id, judul, penulis, penerbit, tahun, rating, eksemplar}
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n > 0 {
		fmt.Println("Buku dengan rating tertinggi:")
		fmt.Println(pustaka[0].judul, pustaka[0].penulis, pustaka[0].penerbit, pustaka[0].tahun)
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].rating < temp.rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Println(pustaka[i].judul, pustaka[i].rating)
	}
}

func CariBuku(pustaka DaftarBuku, n int, rating int) {
	left, right := 0, n-1
	found := false
	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].rating == rating {
			fmt.Println("Buku ditemukan:")
			fmt.Println(pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
			found = true
			break
		} else if pustaka[mid].rating < rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&nPustaka)
	DaftarkanBuku(&pustaka, &nPustaka)
	UrutBuku(&pustaka, nPustaka)
	CetakTerfavorit(pustaka, nPustaka)
	Cetak5Terbaru(pustaka, nPustaka)
	var rating int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, nPustaka, rating)
}
