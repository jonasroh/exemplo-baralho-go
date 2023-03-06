package main

import "fmt"

// Criando um novo tipo de 'deck'
// que se trata de um slice de strings
type deck []string

func (d deck) print() {
	for i, carta := range d {
		fmt.Println(i, carta)
	}

}
