package main

//Crie um Array de inteiros com 5 elementos e
//imprima cada elemento em uma linha
//separada.

import "fmt"

func main() {
	var numeros = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)
	for _, numeros := range numeros {
		fmt.Println(numeros)
	}

}
