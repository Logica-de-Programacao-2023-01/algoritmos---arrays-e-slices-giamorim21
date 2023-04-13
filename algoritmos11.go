package main

import "fmt"

//Crie um Array bidimensional de inteiros com 2 linhas
//e 3 colunas. Em seguida, solicite ao usuário que informe
//um índice de linha e outro de coluna
//e imprima o valor armazenado nessa posição da matriz.

func main() {
	var matriz = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	var linha, coluna int
	fmt.Println("Digite um índice de linha (0 ou 1):")
	fmt.Scan(&linha)
	fmt.Println("Digite um índice de coluna (0, 1 ou 2):")
	fmt.Scan(&coluna)

	fmt.Printf("O valor armazenado na posição (%d, %d) é %d\n", linha, coluna, matriz[linha][coluna])

}

