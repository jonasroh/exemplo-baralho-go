package main

import (
	"os"
	"testing"
)

func TestNovoDeck(t *testing.T) {
	d := novoDeck()

	if len(d) != 52 {
		t.Errorf("Quantidade de cartas esperadas:52, mas existem: %v", len(d))
	}

	if d[0] != "Às de Espadas" {
		t.Errorf("Primeira carta esperada: Às de Espadas, mas a primeira carta obtida foi: %v", d[0])
	}

	if d[len(d)-1] != "Rei de Ouros" {
		t.Errorf("Última carta esperada: Rei de Ouros, mas a última carta obtida foi: %v", d[len(d)-1])
	}
}

func TestSalvaDeckEImportaDeck(t *testing.T) {
	os.Remove("_decktesting")

	deck := novoDeck()
	deck.salvaEmArquivo("_decktesting")

	deckImportado := importDeck("_decktesting")

	if len(deckImportado) != 52 {
		t.Errorf("Quantidade de cartas esperadas:52, mas existem: %v", len(deckImportado))
	}

	os.Remove("_decktesting")
}
