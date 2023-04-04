package main

import "fmt"

// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice
//e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.

func main() {
	var x int
	fmt.Println("Digite o tamanho do Slice: ")
	fmt.Scan(&x)

	slice := make([]int, x)

	for i := 0; i < x; i++ {
		fmt.Println("Digite o valor do elemento", i+1, ": ")
		fmt.Scan(&slice[i])
	}
	fmt.Println(slice)

	fmt.Println("Slice: ", slice)
	soma := 0
	for _, value := range slice {
		soma += value
	}
	fmt.Println("Soma dos valores: ", soma)

}
