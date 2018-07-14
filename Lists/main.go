package main

import (
	"container/list" // Lista duplamente encadeada
	"fmt"
)

func main() {
	numeros := list.New()
	el := numeros.PushBack(1)
	numeros.PushFront(0)
	numeros.PushBack(2)
	// 012
	for e := numeros.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=")
	numeros.Remove(el)
	for e := numeros.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
