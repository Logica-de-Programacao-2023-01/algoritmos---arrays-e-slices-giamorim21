package main

import "fmt"

//Crie um Array de inteiros com 7 elementos.
//Solicite ao usuário que informe um número que será
//adicionado ao primeiro e ao último elemento do Array.
//Imprima o Array resultante.

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	var numero int
	fmt.Println("Digite um número para adicionar ao primeiro e ao último elemento do array:")
	fmt.Scan(&numero)

	array[0] += numero
	array[6] += numero

	fmt.Println(array)
}

