package main

import (
	"fmt"
	"sync"
	"time"
)

const totalTiket = 50
var konferasiName = "Pemesanan Saady"
var banyakTiket uint = 50 
var pesanan = make([]UserData, 0)

type UserData struct {
	namaDepan string
	namaBelakang string
	email string
	nomorTiket uint
}

var wg = sync.WaitGroup{}

func main() {
	welcome()

	namaDepan,namaBelakang,email,nomorTiket :=  getUserInput()
	validNama,validEmail,validNomorTiket := ValidUserInput(namaDepan,namaBelakang,email,nomorTiket)

	if validEmail && validNama && validNomorTiket{
		pesanTiket(namaDepan,namaBelakang,email,nomorTiket)

		wg.Add(1)

		go kirimTiket(namaDepan,namaBelakang,email,nomorTiket)

		namaDepan :=  GetNamaDepan()
		fmt.Printf("Pesanan Pertama adalah : %v \n" , namaDepan)

		if banyakTiket == 0 {
			fmt.Println("Sayang sekali tiket sudah habis kembali ke surga")
		}
	}else {
			if !validNama{
				fmt.Println("Nama Kamu terlalu pendek")
			}
			if !validEmail {
				fmt.Println("Email anda tidak valid, Silahkan gunakan email yang benar")
			}
			if !validNomorTiket {
				fmt.Println("Nomor yang anda input salah")
			}
		}
	wg.Wait()

}

func welcome() {
	fmt.Printf("Selamat datang di %v Silahkan Pesan. \n", konferasiName)
	fmt.Printf("Kita memiliki %v tiket dan %v tiket yang tersedia . \n", totalTiket, banyakTiket)
	fmt.Printf("Silahkan pesan tiket kamu disini \n")
}

func getUserInput() (string , string , string , uint) {
	var namaDepan string
	var namaBelakang string
	var email string
	var nomorTiket uint

	fmt.Println("Masukkan nama depan anda : ")
	fmt.Scan(&namaDepan)

	fmt.Println("Masukkan nama belakang anda : ")
	fmt.Scan(&namaBelakang)

	fmt.Println("Masukkan email anda : ")
	fmt.Scan(&email)

	fmt.Println("Berapa banyak tiket yang anda pesan : ")
	fmt.Scan(&nomorTiket)	

	return namaDepan,namaBelakang,email,nomorTiket
}

func pesanTiket(namaDepan string,namaBelakang string,email string,nomorTiket uint){
	banyakTiket = banyakTiket - nomorTiket

	var userData = UserData{
		namaDepan: namaDepan,
		namaBelakang: namaBelakang,
		email: email,
		nomorTiket: nomorTiket,
	}

	pesanan = 	append(pesanan, userData)
	fmt.Printf("List Pesanan Anda %v \n ", pesanan)

	fmt.Printf("Terimakasih %v %v sudah memesan tiket sebanyak %v . Kami akan mengirimkan email ke %v \n", namaDepan, namaBelakang,nomorTiket,email)

	fmt.Printf("Sisa tiket %v \n" , banyakTiket)
}

func kirimTiket(namaDepan string,namaBelakang string,email string,nomorTiket uint){
	time.Sleep(10 * time.Second)

	tiket := fmt.Sprintf("%v tiket untuk %v %v", nomorTiket,namaDepan,namaBelakang)
	fmt.Printf("########## \n") 
	fmt.Printf("Mengirim Tiket :\n %v \n Email Pemesan : %v \n", tiket, email)
	fmt.Printf("#########")
	wg.Done()
}

func GetNamaDepan() [] string {
	namaDepan := []string{}
	for _,pesan := range pesanan {
		namaDepan = append(namaDepan,pesan.namaDepan)
	}
	return namaDepan
}