package main

import "fmt"

//Crie um Array de inteiros com 5 elementos.
//Em seguida, crie um novo Slice que contenha apenas os
//elementos do Array que são múltiplos de 3.
//Imprima o novo Slice.

func main() {
	array := [5]int{3, 6, 9, 12, 13}

	slice := []int{}
	for _, valor := range array {
		if valor%3 == 0 {
			slice = append(slice, valor)
		}
	}
	fmt.Println(slice)
}
