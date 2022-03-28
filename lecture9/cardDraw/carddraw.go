package carddraw

import cardgame "ScaleFocus/Lecture8/cardGame"

type dealer interface {
	Deal() *cardgame.Card
	Done() bool
}

func DrawAllCards(dealer dealer) []cardgame.Card {
	var newDeck []cardgame.Card

	/*	currentSlice := make([]cardgame.Card, 0, 52)g
		currentCard := dealer.Deal()

		for currentCard.CardValue != 0 {
			currentSlice = append(currentSlice, *currentCard)
			currentCard = dealer.Deal()
		}
		return currentSlice */
}
