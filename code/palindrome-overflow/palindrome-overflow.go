package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var INTMAX = math.Pow(2, 31)
	const INTMIN = -1 << 33

	result := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		if result > int(INTMAX/10) || (result == int(INTMAX/10) && digit > 7) {
			return 0
		}
		if result < INTMIN/10 || (result == INTMIN/10 && digit < -8) {
			return 0
		}

		result = (result * 10) + digit
	}

	return result
}

func main() {
	tests := []int{123, -456, 120, 0, 1534236469, -2147483412}
	for _, num := range tests {
		fmt.Printf("reverse(%d) = %d\n", num, reverse(num))
	}
}
