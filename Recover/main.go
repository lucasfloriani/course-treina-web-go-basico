package main

import (
	"fmt"
)

func main() {
	fmt.Println("Bem vindo! :)")
	defer func() {
		if ex := recover(); ex != nil {
			fmt.Printf("Desculpe, ocorreu um erro: %s \n", ex)
		}
		fmt.Println("Obrigado por utilizar o nosso software")
	}()
	fmt.Print("Digite um número maior que 10: ")
	var numero int
	fmt.Scanf("%d", &numero)
	if numero <= 10 {
		panic("número é inválido")
	} else {
		fmt.Println("Você digitou um bom número! ;)")
	}
}
