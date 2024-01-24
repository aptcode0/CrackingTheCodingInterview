package main

import "fmt"

func setZeros(matrix [][]int) {
	isRowZero := checkIfRowZero(matrix, 0)
	isColZero := checkIfColZero(matrix, 0)

	recordZeros(matrix)
	setZerosAsRecorded(matrix)

	if isRowZero {
		setZerosInRow(matrix, 0)
	}
	if isColZero {
		setZerosInCol(matrix, 0)
	}
}

func checkIfRowZero(matrix [][]int, row int) bool {
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[row][i] == 0 {
			return true
		}
	}
	return false
}

func checkIfColZero(matrix [][]int, col int) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][col] == 0 {
			return true
		}
	}
	return false
}

func recordZeros(matrix [][]int) {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
}

func setZerosAsRecorded(matrix [][]int) {
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
}

func setZerosInRow(matrix [][]int, row int) {
	for i := 0; i < len(matrix[0]); i++ {
		matrix[row][i] = 0
	}
}

func setZerosInCol(matrix [][]int, col int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][col] = 0
	}
}

func main() {
	matrix := [][]int{
		{1, 0, 1, 1},
		{2, 2, 2, 2},
		{0, 3, 3, 3},
		{4, 4, 4, 4},
	}
	setZeros(matrix)
	display(matrix)
}

func display(matrix [][]int) {
	for _, r := range matrix {
		fmt.Println(r)
	}
	fmt.Println()
}