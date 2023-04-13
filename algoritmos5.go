package main

import "fmt"

//Crie um Array bidimensional de inteiros com 3 linhas
//e 2 colunas. Solicite ao usuário que informe os valores
//de cada elemento da matriz.
//Em seguida, imprima a matriz resultante.


func main() {
	var matriz [3][2]int
	for linha := range matriz {
		for coluna := range matriz[linha] {
			var valor int
			fmt.Println("Digite um valor:")
			fmt.Scan(&valor)
			matriz[linha][coluna] = valor
		}
	}
	for _, linha := range matriz {
		fmt.Printf("%d\n", linha)
	}
}
