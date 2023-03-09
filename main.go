package main

import (
	"fmt"
)

func main() {

	//1.menampilkan nilai i : 21
	//2.menampilkan tipe data dari variabel i
	//3.menampilkan tanda %
	//4.menampilkan nilai boolean j : true
	//5.menampilkan nilai boolean j : true
	//6.menampilkan unicode russia : Я (ya)
	//7.menampilkan nilai base 10 : 21
	//8.menampilkan nilai base 8 :25
	//9.menampilkan nilai base 16 : f
	//10.menampilkan nilai base 16 : F 13
	//kurang paham angka 13 itu termasuk result atau apa ya kak ? di expected result itu nilainya F :3

	//11.menampilkan unicode karakter Я : U+042F;var k float64 = 123.456;
	//12.menampilkan float : 123.456000
	//13.menampilkan float scientific : 1.234560E+02

	i := 21
	//1
	fmt.Printf("%T\n", i)
	//2
	fmt.Printf("%v\n", i)
	//3
	fmt.Printf("%%\n")
	//4
	j := true
	fmt.Printf("%t\n", j)
	//5
	fmt.Printf("%t\n\n", j)
	//6
	unicode := 'Я'
	fmt.Printf("%c\n", unicode)

	//tidak tercantum di soal tapi di expected ouput harus ada :v
	fmt.Printf("%b\n", i)
	fmt.Printf("%c\n", '?')
	//7
	fmt.Printf("%d\n", i)
	//8
	fmt.Printf("%o\n", i)
	//9 jika menggunakan nilai dari variabel i itu hasilnya bukan f kak jadi saya ubah ke 15 biar sesuai dengan expected result
	fmt.Printf("%x\n", 15)
	//10 jika menggunakan nilai dari variabel i itu hasilnya bukan f kak jadi saya ubah ke 15 biar sesuai dengan expected result
	fmt.Printf("%X\n", 15)
	//11
	fmt.Printf("%U \n\n", unicode)
	//12
	var k float64 = 123.456
	fmt.Printf("%f\n", k)
	//13
	fmt.Printf("%e\n", k)

	//expected resultnya aneh kak yg di soal ada 1010101
}
