package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateRandomMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(10)
		}
	}
	return matrix
}

func multiplyMatricesParallel(matrix1, matrix2 [][]int) [][]int {
	rows1, cols1 := len(matrix1), len(matrix1[0])
	_, cols2 := len(matrix2), len(matrix2[0])

	result := make([][]int, rows1)
	for i := range result {
		result[i] = make([]int, cols2)
	}

	var wg sync.WaitGroup

	for i := 0; i < rows1; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			for j := 0; j < cols2; j++ {
				for k := 0; k < cols1; k++ {
					result[row][j] += matrix1[row][k] * matrix2[k][j]
				}
			}
		}(i)
	}

	wg.Wait()

	return result
}


func main() {
	rand.Seed(time.Now().UnixNano())

	matrix1 := generateRandomMatrix(1000, 1000)

	matrix2 := generateRandomMatrix(1000, 1000)

	start := time.Now()
_ = multiplyMatricesParallel(matrix1, matrix2)
	elapsed := time.Since(start)
	fmt.Print(elapsed)
}
