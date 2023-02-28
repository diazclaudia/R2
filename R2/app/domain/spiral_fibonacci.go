package domain

import "errors"

func SpiralFibonacci(cols, rows int) ([][]int, error) {
	var matrix = make([][]int, rows)

	if cols < 0 || rows < 0 {
		return matrix, errors.New("not valid rows and cols")
	}
	for a := 0; a < rows; a++ {
		matrix[a] = make([]int, cols)
	}

	top := 0
	bottom := rows - 1
	left := 0
	right := cols - 1
	dir := 0
	counter := 0
	fibonacciNumber := fibonacci(rows * cols)

	for top <= bottom && left <= right {
		if dir == 0 {
			for i := left; i < right+1; i++ {
				matrix[top][i] = fibonacciNumber[counter]
				counter++
			}
			top++
			dir = 1
		} else if dir == 1 {
			for i := top; i < bottom+1; i++ {
				matrix[i][right] = fibonacciNumber[counter]
				counter++
			}
			right--
			dir = 2
		} else if dir == 2 {
			for i := right; i > left-1; i-- {
				matrix[bottom][i] = fibonacciNumber[counter]
				counter++
			}
			bottom--
			dir = 3
		} else if dir == 3 {
			for i := bottom; i > top-1; i-- {
				matrix[i][left] = fibonacciNumber[counter]
				counter++
			}
			left++
			dir = 0
		}
	}

	return matrix, nil
}

func fibonacci(lim int) []int {
	var aux, fib = 0, 1
	fibonacciArray := make([]int, lim)

	for index := 0; index < lim; index++ {
		aux += fib
		fib = aux - fib
		fibonacciArray[index] = fib
	}
	return fibonacciArray
}
