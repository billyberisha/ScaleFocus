package carddraw

import cardgame "ScaleFocus/Lecture8/cardGame"

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {
	currentSlice := make([]cardgame.Card, 0, 52)g
	currentCard := dealer.Deal()

	for currentCard.CardValue != 0 {
		currentSlice = append(currentSlice, *currentCard)
		currentCard = dealer.Deal()
	}
	return currentSlice
}
