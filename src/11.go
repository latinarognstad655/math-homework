
package main

import "math/rand"

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	z := x + y
	fmt.Println("The sum of", x, "and", y, "is", z)
}