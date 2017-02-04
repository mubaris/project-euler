package main

import "fmt"

func sumSquareDiff(n int64) int64 {
	return (3*n*n*n*n + 2*n*n*n - (3 * n * n) - (2 * n)) / 12
}

func main() {
	fmt.Println(sumSquareDiff(100))
}
