package main

func main() {
	cartas := importDeck("minhas_cartas")
	cartas.embaralhaDeck()
	cartas.print()
}
