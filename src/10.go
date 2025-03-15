package main

import "math/rand"

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Println("The result of x + y is:", x+y)
}
