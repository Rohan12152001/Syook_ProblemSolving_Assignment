// Easy Question 1
package main

import (
	"fmt"
	"math"
)

// Check for the factors of the given number starting from 2 upto the
// sqaure root of the given number & go on adding up the factors
// We also use the fact that Factors are always in pairs
//{Time complexity: O(n^(1/2)}
func PerfectNumber(number float64) string{
	factorSum := 1
	var IntNumber int = int(number)

	var iterator float64
	var IntIterator int
	for iterator=2; iterator<math.Sqrt(number);iterator++{
		IntIterator = int(iterator)
		if IntNumber%IntIterator==0{
			tmp := IntNumber/IntIterator
			factorSum+=IntIterator
			factorSum+=tmp
		}
	}

	if factorSum==IntNumber{
		return "Perfect"
	}else if factorSum<IntNumber{
		return "Deficient"
	}else{
		return "Abundant"
	}
}


// Input is always one number (n)
func main() {
	// Common variables
	var num float64

	// Input
	fmt.Println("Enter the Number: ")
	fmt.Scanln(&num)

	// feed the number
	response := PerfectNumber(num)
	fmt.Printf("Output: %s", response)
}