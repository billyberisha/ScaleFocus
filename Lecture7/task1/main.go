package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	cardValue int
	cardSuit  int
}

type Deck struct {
	deck []Card
}

func newDeck() *Deck {
	cardValue := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	cardSuit := []int{1, 2, 3, 4}

	deck := Deck{make([]Card, 0, 52)}

	for v := range cardValue {
		for s := range cardSuit {
			currentCard := Card{
				cardValue: v,
				cardSuit:  s,
			}
			deck.deck = append(deck.deck, currentCard)
		}
	}
	return &deck
}

func (d Deck) Shuffle() *Deck {
	for i := 0; i < 100; i++ {

		randomNum1 := random()
		randomNum2 := random2()
		var oldCard Card = d.deck[randomNum1]
		d.deck[randomNum1] = d.deck[randomNum2]
		d.deck[randomNum2] = oldCard

	}
	return &d
}

func random() int {
	rand.Seed(time.Now().Local().UnixMicro())
	return rand.Intn(52)
}

func random2() int {
	rand.Seed(time.Now().Local().UnixMilli())
	return rand.Intn(52)
}

func (d *Deck) Deal() *Card {
	if len(d.deck) == 0 {
		return &Card{}
	}

	currentCard := d.deck[0]
	d.deck = d.deck[1:]
	return &currentCard
}

func main() {
	d := newDeck()
	d.Shuffle()
	for i := 0; i < 55; i++ {
		fmt.Println(*d.Deal())
	}
}
