package main

import "fmt"

//Crie um Array de floats com 10 elementos.
//Crie um novo Slice que contenha apenas os elementos do
//Array que são maiores que 5.
//Imprima o novo Slice.

func main() {

	array := [10]float64{3.2, 5.5, 7.6, 9.1, 2.3, 1.5, 6.8, 4.2, 8.0, 5.9}
	fmt.Println("Essa é o array:", array)

	var slice []float64
	for _, valor := range array {
		if valor > 5 {
			slice = append(slice, valor)
		}
	}

