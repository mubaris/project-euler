package main

import "fmt"

func largestPrimeFactor(n int64) int64 {
	var lpf int64 = 2
	for n > lpf {
		if n%lpf == 0 {
			n = n / lpf
			lpf = 2
		} else {
			lpf += 1
		}
	}
	return lpf
}

func main() {
	fmt.Println(largestPrimeFactor(600851475143))
}
