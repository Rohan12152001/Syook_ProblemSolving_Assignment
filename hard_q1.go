// Hard Question No.1
package main

import (
	"fmt"
	_ "math"
)

// While iterating if the respective number is divisible by the (k or l or m or n) then mark it as true
// And while doing this we keep a counter for "true"
func countFish(k,l,m,n,total int) (int, []int){
	var uncaughtFishes []int
	counter := 0

	for number:=1;number<=total;number++{
		if number%k==0 || number%l==0 || number%m==0 || number%n==0 {
			counter+=1
		}else{
			uncaughtFishes = append(uncaughtFishes, number)
		}
	}

	return counter, uncaughtFishes
}

// Input is always one number (n)
func main() {
	// Common variables
	variableCount := 5

	// Input
	input := make([]int, variableCount)			// We take k,l,m,n,total
	for index, _ := range input {
		fmt.Scanf("%d", &input[index])
	}

	// Pass the matrix
	response, uncaughtFishes := countFish(input[0], input[1], input[2], input[3], input[4])
	fmt.Printf("Number of fishes caught: %d\n", response)
	fmt.Printf("The list of fishes uncaught: %v", uncaughtFishes)
}