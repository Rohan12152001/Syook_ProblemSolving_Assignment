// Easy Question 3
package main

import (
	"fmt"
)

// We can optimise more
// if a number can be expressed as a power of 2 then it has a direct answer
// For eg: 8 = 2^3 (Here the function will ouput 3)
// So likewise when we have this type of number we can save calculations
func CountTrailsToOne(number int) int{
	count := 0

	for number!=1{
		count+=1
		// number is odd
		if number%2==1{
			number = number*3 + 1
		}else{
			// for even
			number/=2
		}
	}

	return count
}

// Input is always one number (n)
func main() {
	// Common variables
	var num int

	// Input
	fmt.Println("Enter the Number: ")
	fmt.Scanln(&num)

	// feed the number
	response := CountTrailsToOne(num)
	fmt.Printf("Steps required: %v", response)
}