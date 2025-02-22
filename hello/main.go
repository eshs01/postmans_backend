package main

import (
	"fmt"
	"math"
)

const b string = "suii"

func main() {
	fmt.Println(b)
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
