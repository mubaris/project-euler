package main

import "fmt"

func main() {
	for i := 2520; i >= 0; i++ {
		if i%2 == 0 && i%3 == 0 && i%5 == 0 && i%7 == 0 && i%9 == 0 && i%11 == 0 && i%12 == 0 && i%13 == 0 && i%14 == 0 && i%15 == 0 && i%16 == 0 && i%17 == 0 && i%18 == 0 && i%19 == 0 && i%20 == 0 {
			fmt.Println(i)
			break
		}
	}
}
