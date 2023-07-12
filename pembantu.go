package main

import "strings"

func ValidUserInput(namaDepan string,namaBelakang string, email string, nomorTiket uint)(bool, bool, bool){

	validNama := len(namaDepan) >= 2 && len(namaBelakang) >=2
	validEmail := strings.Contains(email,"@")
	validNomorTiket := nomorTiket > 0 && nomorTiket < banyakTiket
	return validNama, validEmail, validNomorTiket
}
