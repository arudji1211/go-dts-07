package materi

import "fmt"

const (
	key1 = "key1"
	key2 = "key2"
)

func Penjelasan() {
	// inisialisasi array
	// param pertama itu tipe data
	// param ke dua itu jumlah data
	arrOfInt := make([]int, 5)
	fmt.Println("Array of integer : ", arrOfInt)

	// inisialisasi slice
	sliceOfInt := []int{}
	fmt.Println("Slice of integer : ", sliceOfInt)

	// copy by value
	copySliceOfInt := []int{}
	copy(copySliceOfInt, sliceOfInt)

	//copy by reference
	//copySliceOfInts := sliceOfInt

	//edit value array index ke 2
	arrOfInt[2] = 100

	// inisialisasi map
	mapOfString := map[string]string{}
	mapOfString[key1] = "aruji"
	mapOfString[key2] = "hermatyar"
	fmt.Println("sebelum delete ", mapOfString)

	//delete key
	delete(mapOfString, key1)
	fmt.Println("setelah delete ", mapOfString)

	i := 0
	for i < 11 {
		i++
		fmt.Println("nilai i : ", i)
	}

	//for range
	for idx, val := range arrOfInt {
		fmt.Printf("nilai idx : %v, nilai value : %v\n", idx, val)
	}
}
