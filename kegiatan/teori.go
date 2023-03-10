package kegiatan

import (
	"fmt"
)

//Materi function start from line 8 - 101
//Materi struct start from line 104 - 116
//Materi Init Function from line 118 - 120

type (
	Operations     string
	OperationsFunc func(num []int) int
)

const (
	TAMBAH Operations = "tambah"
	BAGI   Operations = "bagi"
	KURANG Operations = "kurang"
	KALI   Operations = "kali"
)

// single return
func add(num1, num2 int) int {
	return num1 + num2
}

// multiple return
func multiply(num1, num2, pengali int) (int, int) {
	return num1 * pengali, num2 * pengali
}

// multiple var
func multiplymany(pengali int, nums ...int) []int {
	var result []int
	for _, bils := range nums {
		result = append(result, bils*pengali)
	}
	return result
}

// function return function
func generateOps(ops Operations) OperationsFunc {
	var fn OperationsFunc
	switch ops {
	case TAMBAH:
		fn = func(nums []int) int {
			res := 0
			for _, bils := range nums {
				if res == 0 {
					res = bils
				} else {
					res += bils
				}

			}
			return res
		}
	case KURANG:
		fn = func(nums []int) int {
			res := 0
			for _, bils := range nums {
				if res == 0 {
					res = bils
				} else {
					res -= bils
				}
			}
			return res
		}
	case BAGI:
		fn = func(nums []int) int {
			res := 0
			for _, bils := range nums {
				if res == 0 {
					res = bils
				} else {
					res /= bils
				}
			}
			return res
		}
	case KALI:
		fn = func(nums []int) int {
			res := 0
			for _, bils := range nums {
				if res == 0 {
					res = bils
				} else {
					res *= bils
				}
			}
			return res
		}
	}
	return fn
}

// function have function as parameter and execute this params function
// then return of the return value of params function
func ExecOps(fn OperationsFunc, nums ...int) int {
	result := fn(nums)
	return result
}

type User struct {
	Id    int
	Nama  string
	Email string
}

func (usr User) Greeting() {
	fmt.Printf("Halo %s, id kamu adalah %d\n", usr.Nama, usr.Id)
}

//func init() {
//	fmt.Println("ini init function")
//}

func Materi() {
	fmt.Printf("Materi Function \n")
	fmt.Printf("----------------------\n\n")
	res := add(1, 2)
	fmt.Println("1 + 2 adalah ", res)

	res1, res2 := multiply(2, 3, 9)
	fmt.Println("multiply 1 adalah ", res1)
	fmt.Println("multiply 2 adalah ", res2)

	fn := generateOps(KALI)

	result := ExecOps(fn, 2, 3, 4, 5)

	fmt.Println(result)
	fmt.Printf("----------------------\n\n")
	///materi struct
	fmt.Printf("Materi Struct \n")
	fmt.Printf("----------------------\n\n")
	//1.mengisi struct
	arudji := User{
		Id:    1,
		Nama:  "arudji",
		Email: "bantalku890@gmail.com",
	}

	//2.mengakses struct
	fmt.Printf("ID : %d \n", arudji.Id)
	fmt.Printf("Nama : %v \n", arudji.Nama)
	fmt.Printf("Email : %v \n", arudji.Email)

	arudji.Greeting()

	///materi init
	fmt.Printf("\n\nMateri init function \n")
	fmt.Printf("----------------------\n\n")
	///cara membuat init function yaitu func init()
}
