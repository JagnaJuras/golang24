package main

import (
	"fmt"
	"math/rand"
)

func monty_hall(liczba_rund int, strategia_gracza int) {
	Wygrane := 0
	for _ = range liczba_rund {
		Wybór := rand.Intn(3) // losowo wybiera liczbę od 0 do 2
		Otwarcie := rand.Intn(3)
		for Otwarcie == Wybór || Otwarcie == 0 {
			Otwarcie = rand.Intn(3)
		}
		if strategia_gracza == 1 {
			Wybór_2 := rand.Intn(3)
			for Wybór_2 == Wybór || Wybór_2 == Otwarcie {
				Wybór_2 = rand.Intn(3)
			}
			Wybór = Wybór_2

		}
		if Wybór == 0 {
			Wygrane++
		}
	}
	fmt.Println(Wygrane)
}

func main() {
	monty_hall(100, 0)
	monty_hall(100, 1)
}
