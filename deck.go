package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/shogo82148/go-shuffle"
)

// Criando um novo tipo de 'deck'
// que se trata de um slice de strings
type deck []string

func novoDeck() deck {
	cartas := deck{}

	naipesCarta := []string{"Espadas", "Paus", "Copas", "Ouros"}
	valoresCarta := []string{"Às", "Dois", "Três", "Quatro", "Cinco", "Seis", "Sete", "Oito", "Nove", "Dez", "Valete", "Dama", "Rei"}

	for _, naipe := range naipesCarta {
		for _, valor := range valoresCarta {
			cartas = append(cartas, valor+" de "+naipe)
		}
	}
	return cartas
}

func (d deck) print() {
	for i, carta := range d {
		fmt.Println(i, carta)
	}
}

func embaralhar(d deck, tamanhoMao int) (deck, deck) {
	return d[:tamanhoMao], d[tamanhoMao:]
}

func (d deck) transformaEmString() string {
	return strings.Join(d, ",")
}

// func para salvar o baralho em um arquivo
func (d deck) salvaEmArquivo(nomeArquivo string) error {
	return ioutil.WriteFile(nomeArquivo, []byte(d.transformaEmString()), 0666)
}

func importDeck(nomeArquivo string) deck {
	bs, err := ioutil.ReadFile(nomeArquivo)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) embaralhaDeck() {
	shuffle.Slice(d)
}
