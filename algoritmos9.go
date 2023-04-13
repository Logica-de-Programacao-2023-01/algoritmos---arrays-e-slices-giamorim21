package main

import "fmt"

//Crie um Array de floats com 6 elementos.
//Solicite ao usuário que informe um número que será
//adicionado em todas as posições do Array.
//Imprima o Array resultante.

func main() {
	var n [6]float64
	var valor float64
	fmt.Println("Informe um número a ser adicionado ao array:")
	fmt.Scan(&valor)

	for i := 0; i < len(n); i++ {
		n[i] = valor
	}
	fmt.Println(n)
}
