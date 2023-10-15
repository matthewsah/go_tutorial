package main

import "fmt"

func main() {
	// arrays in go have a fixed size

	nums := [3]int{1, 5, 3}

	fmt.Println(nums)

	// usually we will use slices, not arrays
	// you can make a slice using make([] type, len, cap (optional))
	// if the approximate size of a slice is known before it is filled up,
	// the program can be made faster by creating the slice with that size ahead of time
	nums_slice := make([]int, 3, 6)
	nums_slice[0] = 13
	fmt.Println(nums_slice)

	// you can also make a slice with a specific set of values using
	// []type{val1, val2, val3}
	mySlice := []string{"hello", "this", "is", "a", "slice"}
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))

	// 2d slices
	// matrix := make([][]int, 9)
	// rows, cols := 9, 9
	// for i := 0; i < rows; i++ {
	// 	row := make([]int, 9)
	// 	for j := 0; j < cols; j++ {
	// 		row[j] = i * j
	// 	}
	// 	matrix[i] = row
	// }
	matrix := make([][]int, 0)
	rows, cols := 9, 9
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}

	for i := 0; i < rows; i++ {
		fmt.Println(matrix[i])
	}

	msg := []string{"hello", "this", "is", "a", "slice"}
	for i, word := range msg {
		fmt.Println(i, word)
	}
}
