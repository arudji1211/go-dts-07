package materi

import "fmt"

func Quiz() {
	//[1,1,2,5,4,7,3,4,4,6,5,7,9,9]
	arOfInt := [14]int{1, 1, 2, 5, 4, 7, 3, 4, 4, 6, 5, 7, 9, 9}
	arOfDuplicate := [10]int{}

	for _, val := range arOfInt {
		///
		arOfDuplicate[val] += 1
	}

	data := []int{}
	for i := 0; i < len(arOfDuplicate); i++ {
		if arOfDuplicate[i] > 1 {
			data = append(data, i)
		}
	}
	fmt.Println(data)
}
