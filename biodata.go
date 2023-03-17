package main

import (
	"os"
	"strconv"

	"github.com/arudji1211/go-dts-07/quiz"
)

type Murid map[string]string

func main() {

	//data murid
	Murids := []quiz.MapMurid{
		0: {"Nama": "Arudji Hermatyar", "Alamat": "Makassar", "Pekerjaan": "Mahasiswa", "Alasan": "Upgrade ilmu"},
		1: {"Nama": "Rachmat Hidayat", "Alamat": "Makassar", "Pekerjaan": "Mahasiswa", "Alasan": "Persiapan sebelum lulus"},
		2: {"Nama": "Fajrin Ramadhan", "Alamat": "Makassar", "Pekerjaan": "Mahasiswa", "Alasan": "Upgrade ilmu"},
		3: {"Nama": "Muh Faqih", "Alamat": "Makassar", "Pekerjaan": "Mahasiswa", "Alasan": "Tertarik dengan microservices"},
		4: {"Nama": "Faiz", "Alamat": "Makassar", "Pekerjaan": "Mahasiswa", "Alasan": "Persiapan sebelum lulus"},
	}
	//inisiasi
	MuridFunc := quiz.NewMurid(Murids)
	//tangkap data argumen
	var argsRaw = os.Args
	//konversi tipe data ke int
	id, err := strconv.Atoi(argsRaw[1])
	if err != nil {
		panic("Argumen harus berupa angka")
	}
	//ambil data user
	user := MuridFunc.GetMurid(int(id))
	///show data user
	MuridFunc.ShowMurid(user)
}
