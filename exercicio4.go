package main

import "fmt"

//Crie um Array de floats com 6 elementos e
//calcule a média dos valores armazenados no Array.

func main() {
	var numeros = [6]float64{1, 2, 3, 4, 5, 6}
	var soma float64
	soma = 0
	for _, num := range numeros {
		soma = soma + num
	}
	media := soma / 6
	fmt.Println(media)
}
