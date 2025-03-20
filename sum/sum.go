package main

import "fmt"

func main() {
	a:=12
	b:=24
	fmt.Println(sum(a, 8))
	fmt.Println(mult(2, b))
}

func sum(a, b int) int {
	return a + b
}
func mult(a,b int)int{
	return a*b
}