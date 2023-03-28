package main

//Crie um Array de floats com 4 elementos
//e calcule o produto dos valores armazenados no Array.

import "fmt"

func main() {
	var numeros = [4]float64{1, 3, 5, 7}
	var produto float64
	produto = 1
	for _, num := range numeros {
		produto = produto * num
	}
	fmt.Println(produto)
}
