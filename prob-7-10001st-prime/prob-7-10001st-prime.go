package main

import "fmt"

func isPrime(n int64) int {
	for i := int64(2); i*i <= n; i++ {
		if n%i == int64(0) {
			return 0
		}
	}
	return 1
}

func prime(n int64) int64 {
	var counter int64 = 0
	if n == 1 {
		return 2
	}
	for i := int64(3); ; i += 2 {
		if isPrime(i) == 1 {
			counter++
			if counter == n-1 {
				return i
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(prime(10001))
}
