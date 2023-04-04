package main

import "fmt"

func main() {
	var n int
	var primos []int

	fmt.Println("Digite o valor de n:")
	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		var primo bool = true
		for x := 2; x < i; x++ {
			if i%x == 0 {
				primo = false
				break
			}
		}
		if primo {
			primos = append(primos, i)
		}
	}
	fmt.Println(primos)
}
