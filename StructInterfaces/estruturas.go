package main

import "errors"

// Acelerador : interface que define a possibilidade de acelerar
type Acelerador interface {
	Acelerar() error
}

// Freador : interface que define a possibilidade de frear
type Freador interface {
	Frear() error
}

// Carro : struct que define um carro
type Carro struct {
	Modelo     string
	Marca      string
	Velocidade float32
}

func (carro *Carro) Acelerar() error {
	if carro.Velocidade < 15 {
		carro.Velocidade += 5
		return nil
	}
	return errors.New("o carro já está em sua velocidade máxima")
}

func (carro *Carro) Frear() error {
	if carro.Velocidade > 0 {
		carro.Velocidade -= 5
		return nil
	}
	return errors.New("o carro já está parado")
}
