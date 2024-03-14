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

		// "otwarcie" wybranej liczby drzwi (odejmuje
		dostepne_drzwi := n - liczba_otwarc
		if wybor >= n-liczba_otwarc {
			dostepne_drzwi--
		}

		// zmienia wybór, jeśli strategia gracza polega na zmianie
		if strategia_gracza == 1 {
			nowy_wybor := rand.Intn(dostepne_drzwi)
			for nowy_wybor == wybor {
				nowy_wybor = rand.Intn(dostepne_drzwi)
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
