package handlers

import "fmt"

func Fibonacci(row, col int) {
	dataLen := row * col
	fibonacci := []int{}

	// cari data fibonaci sebanyak dataLen
	fibonacci = append(fibonacci, 0, 1)
	for i := 2; i < dataLen; i++ {
		f := fibonacci[i-1] + fibonacci[i-2]
		fibonacci = append(fibonacci, f)
	}

	fmt.Println(fibonacci)

	// masukkan ke dalam result [][]int
	result := map[int][]int{}
	r := 0
	c := 0

	for _, v := range fibonacci {
		result[r] = append(result[r], v)
		if c < col-1 {
			c++
		} else {
			c = 0
			r++
		}
	}

	fmt.Println(result)
}
