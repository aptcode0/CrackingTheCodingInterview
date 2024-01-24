package main

import "fmt"

/*
Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4
bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
*/
func rotate(matrix [][]int) {
	N := len(matrix)

	for lvl := 0; lvl < N/2; lvl++ {
		start, end := lvl, N-lvl-1

		for i := start; i < end; i++ {
			// top > row constant, col start to end
			tmp := matrix[start][i]

			// top = left (col constant, row end to start)
			matrix[start][i] = matrix[end-i+start][start]

			// left = bottom (row constant, col end to start)
			matrix[end-i+start][start] = matrix[end][end-i+start]

			//bottom = right (col constant, row start to end)
			matrix[end][end-i+start] = matrix[i][end]

			// right = top
			matrix[i][end] = tmp
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	display(matrix)

	matrix = [][]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{3, 3, 3, 3},
		{4, 4, 4, 4},
	}
	rotate(matrix)
	display(matrix)
}

func display(matrix [][]int) {
	for _, r := range matrix {
		fmt.Println(r)
	}
	fmt.Println()
}