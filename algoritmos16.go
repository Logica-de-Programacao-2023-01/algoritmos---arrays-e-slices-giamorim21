package main

import "fmt"

//Crie um Array de inteiros com 10 elementos.
//Crie um novo Slice que contenha apenas os elementos
//pares do Array original.
//Imprima o novo Slice.

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var slice []int
	for _, valor := range array {
		if valor%2 == 0 {
			slice = append(slice, valor)
		}
	}
