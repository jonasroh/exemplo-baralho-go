package main

func main() {
	// criando um slice
	cartas := deck{"Rei de Espadas", novaCarta()}

	//adicionando um novo elemento
	cartas = append(cartas, "Três de Ouros")

	cartas.print()
}

func novaCarta() string {
	return "Cinco de Copas"
}
