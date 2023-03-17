package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type dataDynamic interface{}

func main() {
	var data1 dataDynamic = []string{"coba1", "coba2", "coba3", "coba4"}
	var data2 dataDynamic = []string{"bisa1", "bisa2", "bisa3", "bisa4"}

	genChan := GenerateChan(8, []dataDynamic{data1, data2})
	//
	fmt.Println("====================================================")
	fmt.Println("Concurrency without mutex")
	for a := 0; a < 2; a++ {
		PoolingFunc(genChan)
	}
	fmt.Println("Jumlah GOROUTINE Berjalan : ", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
	fmt.Println("===================================================")
	fmt.Println("Concurency with mutex")
	var mutex sync.Mutex
	PoolingFuncSync1(8, &mutex, data1, data2)
	fmt.Println("Jumlah GO ROUTINE Berjalan : ", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
}

func GenerateChan(totalData int, data []dataDynamic) <-chan dataDynamic {
	chanData := make(chan dataDynamic, totalData)

	go func() {
		defer close(chanData)
		for i := 1; i <= totalData; i++ {
			if i%2 == 0 {
				chanData <- data[0]
			} else {
				chanData <- data[1]
			}
		}
	}()
	return chanData
}

func PoolingFunc(chanInterface <-chan dataDynamic) {
	go func() {
		for i := range chanInterface {
			OutText(i)
		}
	}()
}

// func PoolingFuncSync(Iterator int, data1, data2 dataDynamic) {
// 	var mutex sync.Mutex
// 	data := data1
// 	for i := 1; i <= Iterator; i++ {
// 		go func() {
// 			mutex.Lock()
// 			OutText(data)
// 			var a []string = data.([]string)
// 			var b []string = data1.([]string)
// 			if a[0] == b[0] {
// 				data = data2
// 			} else {
// 				data = data1
// 			}
// 			mutex.Unlock()
// 		}()

// 	}
// }

func PoolingFuncSync1(iterator int, mutex *sync.Mutex, data1, data2 dataDynamic) {
	var data dataDynamic = data1
	for i := 0; i < 2; i++ {
		go func() {
			for a := 0; a < 4; a++ {
				mutex.Lock()
				OutText(data)
				var a []string = data.([]string)
				var b []string = data1.([]string)
				if a[0] == b[0] {
					data = data2
				} else {
					data = data1
				}
				mutex.Unlock()
			}
		}()
	}
	go func() {
	}()
}

func OutText(data dataDynamic) {
	fmt.Println(data)
}
