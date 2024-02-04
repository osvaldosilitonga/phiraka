package handlers

import "errors"

func Fibonacci(row, col int) (map[int][]int, error) {
	if row < 1 || col < 1 {
		return nil, errors.New("row or colum must be greater than 0")
	}

	dataLen := row * col
	fibonacci := []int{0, 1}

	// cari data fibonaci sebanyak dataLen
	for i := 2; i < dataLen; i++ {
		f := fibonacci[i-1] + fibonacci[i-2]
		fibonacci = append(fibonacci, f)
	}

	result := map[int][]int{}
	f := 0

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			result[r] = append(result[r], fibonacci[f])
			f++
		}
	}

	return result, nil
}

// Solusi 2
// masukkan ke dalam result [][]int
// r := 0
// c := 0
// for _, v := range fibonacci {
// 	result[r] = append(result[r], v)
// 	if c < col-1 {
// 		c++
// 	} else {
// 		c = 0
// 		r++
// 	}
// }
