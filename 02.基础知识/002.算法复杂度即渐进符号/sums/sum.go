package main

import "fmt"

func sum(n int) int {
	total := 0
	// 从1加到N,1+2+3+4+5+..+N
	for i := 1; i <= n; i++ {
		total = total + i
	}
	return total
}

func main() {
	fmt.Println(sum(100))
}
