package main

import (
	"fmt"
)

func main() {
	/*
		fmt.Println("Olá mundo!")
		fmt.Println("TreinaWeb" + " Cursos")
		fmt.Println("1 + 1 = ", 1+1)
		fmt.Println("1.1 + 2.2 = ", 1.1+2.2)
		fmt.Println(true)

		fmt.Println("INTEIROS SEM SINAL")
		var u1 byte = 255 // == uint8 (0~255)
		fmt.Println(u1)
		var u2 uint16 = 33 // uint16 (0~65.535)
		fmt.Println(u2)
		var u3 uint32 = 44 // uint32 (0~4.294.967.295)
		fmt.Println(u3)
		var u4 uint64 = 55 // uint64 (0~18.446.744.073.790.551.615)
		fmt.Println(u4)

		fmt.Println("INTEIROS COM SINAL")
		var i1 int8 = 127 // int8(-128~127)
		fmt.Println(i1)
		var i2 int16 = 1000 // int16(-32768~32767)
		fmt.Println(i2)
		var i3 rune = 1001 // == int32(-2.147.483.648~2.147.483.647)
		fmt.Println(i3)
		var i4 int64 = 1002 // int64(-9.223.372.036.854.775.808~9.223.372.036.854.775.807)
		fmt.Println(i4)

		var t1 uint = 1010 // troca entre 32 e 64 dependendo da arquitetura do processador
		fmt.Println(t1)
		var t2 int = 2000 // troca entre 32 e 64 dependendo da arquitetura do processador
		fmt.Println(t2)

		fmt.Println("PONTOS FLUTUANTES")
		var f1 float32 = 1
		fmt.Println(f1)
		var f2 float64 = 2
		fmt.Println(f2)
		var f3 complex64 = 3 // 3+0i (i == número imaginário) (igual ao float32 + conjunto de números imaginários)
		fmt.Println(f3)
		var f4 complex128 = 4 // 4+0i (i == número imaginário) (igual ao float64 + conjunto de números imaginários)
		fmt.Println(f4)

		fmt.Println("STRINGS")
		var s1 string = "TreinaWeb Cursos"
		fmt.Println(s1)
		fmt.Println(s1[0])
		var s2 string = `TreinaWeb
		Cursos`
		fmt.Println(s2)

		fmt.Println("BOOLEANOS")
		var b1 bool = true
		fmt.Println(b1)

		fmt.Println("MULTIPLAS DECLARAÇÔES")
		// var nome, sobrenome string = "TreinaWeb", "Cursos"
		var (
			nome      string = "TreinaWeb"
			sobrenome string = "Cursos"
			idade     int    = 10
		)
		fmt.Println(nome, sobrenome)
		fmt.Println(idade)
	*/

	/*
		var nome = "TreinaWeb" // Automáticamente adiciona o tipo
		fmt.Println(nome)
		// var nomeCompleto string = "TreinaWeb Cursos"
		nomeCompleto := "TreinaWeb Cursos" // Atalho
		fmt.Println(nomeCompleto)
	*/

	// const constante string = "TreinaWeb" // Não pode trocar o valor
	const (
		primeiroNome = "TreinaWeb"
		segundoNome  = "Cursos"
	)
	fmt.Println(primeiroNome, segundoNome)
}
