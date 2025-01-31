package main

import "fmt"

func main() {

	x, y := "#", " "

	for i := 0; i < 8; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(x, y)
		}
		fmt.Println()
	}
}
