// Multiples of 3 and 5
// Problem 1

// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func main() {
	var sum int

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)

	// Another way, using an array index and func expression just to be cool
	/*
		var x [1000]int

		fmt.Println(func() int {
			for i, _ := range x {
				if i%3 == 0 || i%5 == 0 {
					x[0] += i
				}
			}
			return x[0]
		}())
	*/
}
