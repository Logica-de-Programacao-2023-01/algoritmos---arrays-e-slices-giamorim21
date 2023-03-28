package main

import "fmt"

//Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5.
//Remova o terceiro elemento e
//imprima o Slice resultante.

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	numeros = append(numeros[:2], numeros[3:]...)
	fmt.Print(numeros)
}
