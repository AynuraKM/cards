package main

import (
	"fmt"
)

func main() {
	deck := Deck{
		cards: []Card{
			{"Туз", "Пики"},
			{"Туз", "Треф"},
			{"Туз", "Буби"},
			{"Туз", "Черви"},
			{"Дама", "Пики"},
			{"Дама", "Треф"},
			{"Дама", "Буби"},
			{"Дама", "Черви"},
			{"Король", "Пики"},
			{"Король", "Треф"},
			{"Король", "Буби"},
			{"Король", "Черви"},
			{"Шестерка", "Пики"},
			{"Шестерка", "Треф"},
			{"Шестерка", "Буби"},
			{"Шестерка", "Черви"},
		},
	}

	fmt.Println(deck)

	deck.shuffle()


	fmt.Println(deck)
}

type Card struct {
	value  string
	colour string
}

type Deck struct {
	cards []Card
}

func (d *Deck) shuffle() {
	n := 3
		arr2 := d.cards[:len(d.cards)/2]
		arr21 := d.cards[len(d.cards)/2:len(d.cards)-n]

		arr3 := d.cards[len(d.cards)-n:len(d.cards)-1]
		arr2 = append(arr2, arr3...)
		arr21 = append(arr2, arr21...)
		d.cards = arr21
}
