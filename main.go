package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}
	str := []byte("CFAEIQAFPGBZO")
	for j := 0; j < 11; j++ {
		if j == 5 {
			for idx, val := range str {
				if idx%2 == 0 {
					fmt.Printf("Character %U %q starts at byte position %d \n", val, val, idx)
				}
			}
		} else {
			fmt.Println("Nilai j = ", j)
		}
	}
}
