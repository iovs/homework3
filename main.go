package main

import "fmt"

func main() {

	var height int
	fmt.Println("\n Задайте размеры шахматной доски")
	fmt.Print("\nВысота: ")
	fmt.Scanf("%d", &height)
	var weight int
	fmt.Print("Ширина: ")
	fmt.Scanf("%d", &weight)
	fmt.Printf("\n Формируем шахматную доску размером %d на %d \n\n", height, weight)

	//size := 8

	for i := 0; i < height; i++ {
		for j := 0; j < weight; j++ {
			if ((i + j) % 2) == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		perenos := "\n"
		fmt.Print(perenos)
		//	break
	}
	fmt.Println("\n Формирование доски закончено. Спасибо.\n ")
}
