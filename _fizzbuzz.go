package main

import (
	"fmt"
)

func fizzbuzz() {
	// for i := 1; i<=100 ; i++ {
	// 	output := ""
	// 	if i % 3 == 0{
	// 		output += "Fizz"
	// 	}
	// 	if i % 5 == 0{
	// 		output += "Buzz"
	// 	}
	// 	if i % 3 != 0 && i % 5 != 0{
	// 		output = fmt.Sprintf("%d", i)
	// 	}
	// 	fmt.Println(output)
	// }

	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
				fmt.Println("FizzBuzz")
		case i%3 == 0:
				fmt.Println("Fizz")
		case i%5 == 0:
				fmt.Println("Buzz")
		default:
				fmt.Println(i)
		}
}
}


func main() {
	fizzbuzz()
}