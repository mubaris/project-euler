package main

import "fmt"

func triplet(n int) int {
	for c := 500; c > 0; c-- {
		for b := 499; b > 0; b-- {
			a := 1000 - c - b
			if (a*a)+(b*b) == c*c {
				return a * b * c
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(triplet(1000))
}
