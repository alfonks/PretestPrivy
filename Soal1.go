package main

import (
	"fmt"
	"strconv"
	"log"
)

func main() {
	fmt.Println(" ")
	//pertama := "19374048" //Assign karakter petama kedalam variable
	//Hapus command pada baris diatas jika ingin menggunakan input dari soal

	//kedua := "aiueobcd" //Assign karakter kedua kedalam variable
	//Hapus command pada baris diatas jika ingin menggunakan input dari soal
	var pertama string
	var kedua string

	fmt.Print("Masukkan string pertama: ")
	fmt.Scanf("%s\n", &pertama)

	fmt.Print("Masukkan string kedua: ")
	fmt.Scanf("%s\n", &kedua)

	indexof7, flagditemukan := 0, 0 //variable ini akan menyimpan index keberapa angka 7 ditemukan

	fmt.Printf("\n------------------------------\n")


	//Diasumsikan bahwa angka 7 yang dicari tidak diketahui letaknya dimana
	//didalam karakter pertama tersebut. Karena urutan karakter tersebut tidak
	//tersorting, maka akan menggunakan sequential search untuk mencari posisi angka ke-7

	//Posisi angka ketujuh berada pada index ke-3
	//Output yang diharapkan, dari posisi angka 7 yaitu index ke-3
	//Akan di print String kedua mulai dari index ke-3 sampai akhir

	fmt.Printf("String pertama = %s\n", pertama)
	fmt.Printf("String kedua = %s\n\n", kedua)

	for v := range pertama { //Fungsi ini akan mencari di index keberapa angka 7 berada

		temp, err := strconv.Atoi(string(pertama[v])) //Mengubah string kedalam bentuk integer yang disimpan kedalam variable temp
		if err != nil { //Saat konversi, jika terdapat error, maka akan mengeluarkan log error
			log.Fatal(err)
		}

		if temp == 7 { //jika temp == 7 maka v akan di assign kedalam  indexof7 dan break dari looping
			indexof7 = v
			flagditemukan = 1
			break
		}
	}

	if (flagditemukan == 1) {
		fmt.Printf("Angka 7 berhasil ditemukan di index ke-%d\n", indexof7)

		if (indexof7 > len(kedua)) {
			fmt.Println("String kedua tidak dapat diprint karena index angka 7 ditemukan lebih besar dari panjang string kedua")
		} else {
			fmt.Printf("String dari index angka 7 ditemukan sampai akhir = %s\n", kedua[indexof7:len(kedua)]) //Mengeprint karakter dari indexof7 sampai akhir
		}

	} else {
		fmt.Printf("Angka 7 tidak ditemukan didalam string pertama\n")
	}

	
}