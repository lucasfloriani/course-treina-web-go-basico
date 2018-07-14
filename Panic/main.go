package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bem-vindo!")
	// defer é executado mesmo quando ocorrer um panic
	defer fmt.Println("Obrigado por utilizar o nosso software")
	fmt.Print("Digite um número maior que 10: ")
	var numero int
	fmt.Scanf("%d", &numero)
	if numero <= 10 {
		// Sai do programa com status 2
		//
		// Quando o programa sai com status 0, isto quer dizer que
		// saiu com sucesso, portanto, caso o programa saia com valor
		// diferente de 0, pode-se dizer que ocorreu um erro.
		//
		// Número de estatus serve para notificar o sistema operacional
		// do que ocorreu.
		panic("Número inválido!")
	} else {
		fmt.Println("Você digitou um bom número! :)")
	}
}
