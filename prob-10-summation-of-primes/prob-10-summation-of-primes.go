package main

import "fmt"

func primeSum(n int64) int64 {
	var sum int64 = 2
	var i int64
	if n < int64(2) {
		return int64(0)
	} else if n == int64(2) {
		return int64(2)
	} else {
		for i = 3; i < n; i += 2 {
			if isPrime(i) == 1 {
				sum += i
			}
		}
	}
	return sum
}

func isPrime(n int64) int {
	var i int64
	for i = 2; i*i <= n; i++ {
		if n%i == 0 {
			return 0
		}
	}
	return 1
}

func main() {
	fmt.Println(primeSum(2000000))
}
