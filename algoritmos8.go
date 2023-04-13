package main

import "fmt"

//Crie um Slice de strings com tamanho 8
//e solicite ao usuário que informe um valor.
//Remova todas as ocorrências desse valor no Slice
//e imprima o resultado.

func main() {
	s := []string{"banana", "uva", "pêssego", "manga", "cereja", "acerola", "maçã", "pera"}

	var valor string
	fmt.Println("Digite um tipo de fruta:")
	fmt.Scan(&valor)

	novoslice := []string{}
	for _, v := range s {
		if v != valor {
			novoslice = append(novoslice, v)
		}
	}
	fmt.Println(novoslice)
}
