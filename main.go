package main

import "fmt"

func PecahKalimat(kalimat string) []string {
	karakter := []string{}
	//pecah
	for i := 0; i < len(kalimat); i++ {
		karakter = append(karakter, string(kalimat[i]))
	}
	return karakter
}

func PrintAndCount(kata []string) map[string]int {
	res := map[string]int{}

	for _, vals := range kata {
		if res[vals] >= 1 {
			res[vals] += 1
		} else {
			res[vals] = 1
		}

		fmt.Printf("%v \n", vals)
	}
	return res
}

func main() {
	//kegiatan.Materi()
	kalimat := "selamat malam"
	kata := PecahKalimat(kalimat)
	counter := PrintAndCount(kata)
	fmt.Println(counter)
}
