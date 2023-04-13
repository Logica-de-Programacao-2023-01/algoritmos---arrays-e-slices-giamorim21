package main

import "fmt"

//Crie um Slice de inteiros com tamanho 10
//e imprima o valor mínimo
//e o valor máximo armazenados no Slice.

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	min := slice[0]
	for _, valor := range slice {
		if valor < min {
			min = valor
		}
	}
	max := slice[0]
	for _, valor := range slice {
		if valor > max {
			max = valor
		}
	}
	fmt.Printf("Valor mínimo: %d\n", min)
	fmt.Printf("Valor máximo: %d\n", max)
}
