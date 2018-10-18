package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func is_even(a int) bool {
	return a % 2 == 0
}

func is_odd(a int) bool {
	return a % 2 == 1
}

func signum(a int) int {
	if a == 0 {
		return 0
	} else if a < 0 {
		return -1
	} else {
		return 1
	}
}

func factorial(a int) int {

	var b int = 1
	for i := 1; i <= a; i++{
		b = b * i
	}
	return b
}

func main() {
	//fmt.Println("Sum of 3 and 4 is: ", sum(3,4))
	//fmt.Println("Is 3 even: ", is_even(3))
	//fmt.Println("Is 4 even: ", is_even(4))
	//fmt.Println("Is 5 odd: ", is_odd(5))
	//fmt.Println("Is 6 odd: ", is_odd(6))
	//fmt.Println(signum(2), signum(9), signum(0), signum(-3))
	//fmt.Println("factorial of 3: ", factorial(3),
	//	           " factorial of 7: ", factorial(7))

	fmt.Println("Hello world!")
}