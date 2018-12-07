/*

A variadic func accepts variable number of input values — zero or more. 
Ellipsis (three-dots) prefix in front of an input type makes a func variadic.

*/

package main 

import "fmt"

//This variadic function takes different data types as arguments
func variadicDifferentTypes(data ...interface{}) {
	fmt.Println("This variadic function takes arguments of different data types.")
	for _, i := range data {
		fmt.Println(i)
	}
}

//This variadic function only takes int  
func variadicSameTypes(numbers ...int) {
	fmt.Println("This variadic function takes integer arguments.")
	for _, i := range numbers {
		fmt.Println(i)
	}
}

//First argument is necessary here
func varidaicFunc(number int, text ...string) {
	fmt.Println("This variadic function has an integer as necessary argument.")
	fmt.Println(number)
	for _, i := range text {
		fmt.Println(i)
	}
}

func main() {
	variadicDifferentTypes("hello", "gophers", "let's rule the world", 1, 2, 3)
	fmt.Printf("\n")

	variadicSameTypes(1, 2, 3, 4, 5)
	fmt.Printf("\n")

	//While using slices make sure you place three dots when passing it as an 
	//argument to variadic function
	arr := []int{10, 11, 12, 13, 14}
	variadicSameTypes(arr...)
	fmt.Printf("\n")

	varidaicFunc(2, "hello", "world")
}
