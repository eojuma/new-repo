package main

import "fmt"

func main() {
	a := 12
	b := 24
	fmt.Println(sum(a, 8))
	fmt.Println(mult(2, b))
	fmt.Println(subt(a, b))
	fmt.Println(div(23.4, 7.5))
	fmt.Println(mod(a, b))

}

func sum(a, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func subt(a, b int) int {
	return a - b
}

func div(a, b float64) float64 {
	return a / b
}

func mod(a, b int) int {
	return (a % b)
	// returns the remainder
}

// The workflow is amazing and by the end of the month I
//  should be able to understand at least 85 per cent of information pertaining Git.

//They say that coding needs critical thinking not fast hands but then i think it neds both critical thinking and fast hands too.
//What do you think?

//The first attempt to anything new is difficult,but then as days go by things get easier,clearer and eventually very familiar.
//Do agree?

///I am almost getting better at merging or rather getting the whole concept of branching.