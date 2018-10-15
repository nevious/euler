/*
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000. 
*/
package main

import "fmt"

/* Alternative:
 * import (
 *  "fmt"
 * )
*/

func main() {
	_cap := 1000
	sum := 0

	for iter :=0 ; iter < _cap ; iter++ {
		if (iter % 3 == 0) || (iter % 5 == 0) {
			sum += iter
			fmt.Printf("Iter: %d\n", iter)
		}
	}

	fmt.Printf("Total Sum: %d\n", sum)
}
