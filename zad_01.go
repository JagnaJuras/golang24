package main

import (
	"fmt"
	"math/rand"
)

func monty_hall(liczba_rund, n, liczba_otwarc int, strategia_gracza int) {
	wygrane := 0
	var nazwa_strategii string

	for _ = range make([]int, liczba_rund) {
		// losuje liczbę jako pierwszy wybór gracza
		wybor := rand.Intn(n)

		// losuje liczby jako pudełka do otwarcia (puste i nie wybrane przez gracza)
		otwarte := make(map[int]bool)
		for len(otwarte) < liczba_otwarc {
			otwarte_pudla := rand.Intn(n)
			if otwarte_pudla != wybor && !otwarte[otwarte_pudla] && otwarte_pudla != 0 {
				otwarte[otwarte_pudla] = true
			}
		}

		// zmienia wybów, jeśli strategia gracza polega na zmianie
		if strategia_gracza == 1 {
			nowy_wybor := rand.Intn(n)
			for nowy_wybor == wybor || otwarte[nowy_wybor] {
				nowy_wybor = rand.Intn(n)
			}
			wybor = nowy_wybor
			nazwa_strategii = "Zmień"
		} else {
			nazwa_strategii = "Nie zmieniaj"
		}

		// sprawdza, czy gracz wygrał
		if wybor == 0 {
			wygrane++
		}
	}

	porażki := liczba_rund - wygrane
	fmt.Printf("Strategia: %s, Wygrane: %d, Porażki: %d\n", nazwa_strategii, wygrane, porażki)
}

func main() {
	monty_hall(100, 3, 1, 0)
	monty_hall(100, 3, 1, 1)
	monty_hall(100, 100, 98, 0)
	monty_hall(100, 100, 98, 1)
}
