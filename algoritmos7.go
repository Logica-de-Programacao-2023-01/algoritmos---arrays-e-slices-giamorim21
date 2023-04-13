package main

import "fmt"

//Crie um Slice de inteiros com o tamanho 5.
//Em seguida, solicite ao usuário que informe
//um número inteiro.
//Adicione esse número ao Slice apenas se ele não estiver
//presente.
//Imprima o Slice resultante.

func main() {
	numeros := make([]int, 5)
	for i := 0; i < len(numeros); {
		var valor int
		fmt.Println("Digite um número inteiro:")
		fmt.Scan(&valor)

		encontrado := false
		for _, num := range numeros {
			if num == valor {
				encontrado = true
				break
			}
		}
		if !encontrado {
			numeros[i] = valor
			i++
		} else {
			fmt.Println("O número já está presente na slice")
		}
	}
	fmt.Println(numeros)
}
