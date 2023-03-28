package main

import "fmt"

//Crie um Array bidimensional de inteiros com 3 linhas
//e 2 colunas. Solicite ao usuário que informe os valores
//de cada elemento da matriz.
//Em seguida, imprima a matriz resultante.

func main() {
	var matriz [3][2]int
	for linha := range matriz {
		for coluna, elemento := range matriz[linha] {
			fmt.Println("Elemento[", linha, "][", coluna, "] =", elemento)
			var x int
			for i := 0; i < x; i++ {
			
			}

		}
	}
}
