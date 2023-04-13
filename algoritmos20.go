package main

//Crie um programa que leia um array de inteiros
//e verifique se ele está ordenado em ordem crescente.

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}

	crescente := true
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			crescente = false
			break
		}
	}

	if crescente {
		fmt.Println("O array está em ordem crescente!")
	} else {
		fmt.Println("O array não está em ordem crescente!")
	}
}
