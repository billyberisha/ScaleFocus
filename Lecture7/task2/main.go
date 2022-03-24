package main

import "fmt"

type Card struct {
	value int
	suit  int
}

func compareCards(cardOne Card, cardTwo Card) (int, bool) {
	if !isValidCard(cardOne) || !isValidCard(cardTwo) {
		return 0, false
	}
	scoreCard1 := cardOne.value*10 + cardOne.suit
	scoreCard2 := cardTwo.value*10 + cardTwo.suit
	if scoreCard1 > scoreCard2 {
		return 1, true
	} else if scoreCard1 < scoreCard2 {
		return -1, true
	}
	return 0, true
}

func maxCard(cards []Card) Card {
	var maxCard Card = cards[0]
	for _, currentCard := range cards {
		if result, ok := compareCards(currentCard, maxCard); ok && result == 1 {
			maxCard = currentCard
		}
	}
	return maxCard
}

func isValidCard(card Card) bool {
	if (card.value >= 2 && card.value <= 15) && (card.suit >= 1 && card.suit <= 4) {
		return true
	}
	return false
}

func main() {
	card1 := Card{value: 4, suit: 2}
	card2 := Card{value: 3, suit: 1}
	card3 := Card{value: 2, suit: 2}
	card4 := Card{value: 1, suit: 1}

	cardSlice := []Card{
		card1, card2, card3, card4,
	}
	if result, ok := compareCards(card1, card2); ok {
		fmt.Println(result)
	} else {
		fmt.Println("something went wrong")
	}
	fmt.Println(maxCard(cardSlice))

}
