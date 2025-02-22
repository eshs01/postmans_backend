package main

import (
	"fmt"
)

const b string = "suii"

func main() {
	i := 10
	for i >= 3 {
		fmt.Println(i)
		i = i - 1
	}
	for k := 0; k != 5; k++ {
		fmt.Println(k)
	}
	for i := range 3 {
		fmt.Println("range", i)
	}
	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
