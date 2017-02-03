package main

import "fmt"

func fibonacci(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {

	sum := 0

	for i := 2; fibonacci(i) <= 4000000; i = i + 3 {
		sum += fibonacci(i)
	}

	fmt.Println(sum)
}
