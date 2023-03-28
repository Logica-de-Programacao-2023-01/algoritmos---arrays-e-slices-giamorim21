package main

import "fmt"

//Crie uma matriz bidimensional de inteiros com 3 linhas e 4 colunas.
//Inicialize cada elemento com o valor da soma do índice da linha e
//o índice da coluna.
//Imprima a matriz resultante.

func main() {
	var matriz [3][4]int

	for linha := range matriz {
		for coluna, elemento := range matriz[linha] {
			elemento = linha + coluna
			fmt.Println("Elemento[", linha, "][", coluna, "]=", elemento)
		}

	}

}
