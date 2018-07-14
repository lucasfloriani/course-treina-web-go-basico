package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estou abrindo a conexão com o banco de dados...")
	defer fmt.Println("Estou fechando a conexão com o banco de dados!")
	defer fmt.Println("Mais um Defer")
	executarConsulta()
}

func executarConsulta() {
	fmt.Println("Estou executando a consulta ao banco de dados...")
}
