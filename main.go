package main

import "fmt"

func main() {

	size := 8

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if ((i + j) % 2) == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		perenos := "\n"
		fmt.Print(perenos)
	}

}
