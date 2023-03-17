package quiz

import "fmt"

type MuridImpl struct {
	data []MapMurid
}

//inisiasi menggunakan konsep polymorphism (interface berubah bentuk menjadi struct)
func NewMurid(dataset []MapMurid) Murid {
	return &MuridImpl{
		data: dataset,
	}
}

//Mengambil data murid berdasarkan ID
func (m *MuridImpl) GetMurid(id int) MapMurid {
	if id < len(m.data) {
		return m.data[id]
	} else {
		temp := MapMurid{"Nama": "Data tidak di temukan", "Alamat": "Data tidak di temukan", "Pekerjaan": "Data tidak di temukan", "Alasan": "Data Tidak di temukan"}
		return temp
	}
}

//Menampilkan data murid ke console
func (m *MuridImpl) ShowMurid(data MapMurid) {
	fmt.Println("Nama		: ", data["Nama"])
	fmt.Println("Alamat		: ", data["Alamat"])
	fmt.Println("Pekerjaan	: ", data["Pekerjaan"])
	fmt.Println("Alasan		: ", data["Alasan"])
}
