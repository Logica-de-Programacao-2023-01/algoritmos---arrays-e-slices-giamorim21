package main

import "fmt"

// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice
//e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.

func main() {
	// Solicita ao usuário o tamanho do Slice
	var x int
	fmt.Print("Digite o tamanho do Slice: ")
	fmt.Scan(&x)

	// Cria o Slice com o tamanho informado pelo usuário
	slice := make([]int, x)

	// Solicita ao usuário os valores dos elementos do Slice
	for i := 0; i < x; i++ {
		fmt.Printf("Digite o valor do elemento %d: ", i+1)
		fmt.Scan(&slice[i])
	}

	// Imprime o Slice e a soma dos valores armazenados
	fmt.Printf("Slice: %v\n", slice)
	sum := 0
	for _, value := range slice {
		sum += value
	}
	fmt.Printf("Soma dos valores: %d\n", sum)
}
