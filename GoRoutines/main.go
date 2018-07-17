package main

import (
	"fmt"
)

func main() {
	var limite int
	// O segundo parâmetro é para criar
	// um channel buferizado (async)
	canal1 := make(chan int, 2000)
	canal2 := make(chan int, 500)
	fmt.Print("Informe um limite: ")
	fmt.Scanf("%d", &limite)

	go conteAte(limite, canal1)
	go func(n int) {
		resultado := 0
		for i := 0; i <= n*10; i++ {
			fmt.Printf(" - [anônimo] O número é %d \n", i)
			resultado = i * 10
		}
		canal2 <- resultado
	}(limite)

	for i := 0; i <= limite; i++ {
		fmt.Printf(" - [main] O número é %d \n", i)
	}

	resultadoCanal1 := <-canal1
	resultadoCanal2 := <-canal2

	fmt.Printf("Canal 1 = %d, Canal 2 = %d \n", resultadoCanal1, resultadoCanal2)
}

func conteAte(limite int, canal1 chan int) {
	resultado := 0
	for i := 0; i <= limite*20; i++ {
		fmt.Printf(" - [conteAte] O número é %d \n", i)
		resultado = i * 20
	}
	canal1 <- resultado
}
