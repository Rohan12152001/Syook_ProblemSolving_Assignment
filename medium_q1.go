// Medium Question 1
package main

import (
	"fmt"
	"math"
)

func countAnswer(matrix [][]int) []int{
	// Maps
	// (the row maps will store the max value for a row & col maps will store least value for thr column)
	var rowMap = map[int]int{}
	var colMap = map[int]int{}

	var array []int

	// Initialise maps
	for row:=0;row< len(matrix);row++{
		rowMap[row] = -1 * math.MaxInt64
	}

	for col:=0;col<len(matrix[0]);col++{
		colMap[col] = math.MaxInt64
	}

	// Feed maps
	for row:=0;row< len(matrix);row++{
		for col:=0;col<len(matrix[0]);col++{
			element := matrix[row][col]
			if element>rowMap[row]{
				rowMap[row] = element
			}
			if element<colMap[col]{
				colMap[col] = element
			}
		}
	}

	// Iterate matrix & form the answer list
	for row:=0;row< len(matrix);row++{
		for col:=0;col<len(matrix[0]);col++{
			element := matrix[row][col]
			if element>=rowMap[row] && element<=colMap[col]{
				array = append(array, element)
			}
		}
	}

	return array
}

// Input is always one number (n)
func main() {
	// Common variables
	var rows int
	var cols int
	var matrix [][]int

	// Input
	fmt.Println("Enter the number of rows: ")
	fmt.Scanln(&rows)

	fmt.Println("Enter the number of columns: ")
	fmt.Scanln(&cols)

	for row:=0;row<rows;row++{
		numbers := make([]int, cols)
		for index, _ := range numbers {
			fmt.Scanf("%d", &numbers[index])
		}
		matrix = append(matrix, numbers)
	}

	// Pass the matrix
	response := countAnswer(matrix)
	fmt.Printf("The numbers are: %v", response)
}