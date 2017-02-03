package main

import . "fmt"

func sumDivisibleBy(n int) int {
	p := 999 / n
	return int(n * p * (p + 1) / 2)
}

func main() {
	total := 0

	/*for i := 1; i < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			total += i
		}
	}*/

	//Optimized Solution

	total = sumDivisibleBy(3) + sumDivisibleBy(5) - sumDivisibleBy(15)

	Println(total)
}
