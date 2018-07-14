package main

import (
	"fmt"
	"strconv" // Usado para conversão de tipos de variaveis
)

func main() {
	// ###### STRING PARA INT ######
	var texto string
	fmt.Print("Digite um número: ")
	fmt.Scanf("%s", &texto)

	// var numero int
	// numero, _ = strconv.Atoi(texto) // Atoi == ASCI To Integer

	// ParseInt(
	//   String para converter,
	//   Conversão para decimal ou binária ou hexadecimal etc,
	//   Variavel vai ser 32 bits ou 64 bits,
	// )
	numero, _ := strconv.ParseInt(texto, 10, 32)
	fmt.Println(numero)

	// ###### INT PARA FLOAT ######
	var conv float64 = float64(numero) // Usado quando um tipo pode estar dentro de outro (int como float)
	fmt.Println(conv)
}
