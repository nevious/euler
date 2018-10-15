package main

import "fmt"
import "math"

func main() {
	_cap := 4 * int(math.Pow(10, 6))
	fmt.Printf("Cap: %d\n", _cap)
	fibos := []int{1, 2}
	sum := 0
	new_val := 0

	for new_val < _cap {
		sl_length := len(fibos)
		new_val = fibos[sl_length-2] + fibos[sl_length-1]

		if (new_val > _cap){
			break
		}

		fibos = append(fibos, new_val)
		fmt.Printf("F: %d\n", fibos)
	}

	for _, value := range fibos { // _ is the index
		if (value % 2 == 0){
			sum += value
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}
