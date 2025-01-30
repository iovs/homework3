package main

import "fmt"

func main() {
	var s string
	const a = " "
	const b = "#"

	s = ` # # # #
# # # #
 # # # #
# # # #`

	println(s) // Вариант 1

	fmt.Println(a, b, a, b, a, b, a, b) // Вариант 2
	fmt.Println(b, a, b, a, b, a, b, a)
	fmt.Println(a, b, a, b, a, b, a, b)
	fmt.Println(b, a, b, a, b, a, b, a)

}
