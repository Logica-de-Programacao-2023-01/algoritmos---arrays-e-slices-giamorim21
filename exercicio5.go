package main

import "fmt"

//Crie uma matriz bidimensional de inteiros com 3 linhas
//e 4 colunas.
//Inicialize cada elemento com o valor da soma do índice
//da linha e
//o índice da coluna.
//Imprima a matriz resultante.

func main() {
	// cria a matriz de inteiros
	matriz := [3][4]int{}
	
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			matriz[i][j] = i + j
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println()
	}
}
