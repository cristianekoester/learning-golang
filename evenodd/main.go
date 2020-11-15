package main

import "fmt"

func main() {
	numbers := make([]int, 11)
	for n := range numbers {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}
}
