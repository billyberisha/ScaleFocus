package cardgame

type Card struct {
	CardValue int
	CardSuit  int
}

type Deck struct {
	Deck []Card
}

func NewDeck() *Deck {
	cardValue := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	cardSuit := []int{1, 2, 3, 4}

	deck := Deck{make([]Card, 0, 52)}

	for _, v := range cardValue {
		for _, s := range cardSuit {
			currentCard := Card{
				CardValue: v,
				CardSuit:  s,
			}
			deck.Deck = append(deck.Deck, currentCard)
		}
	}
	return &deck
}

func (d *Deck) Deal() (*Card, error) {
	if len(d.Deck) != 0 {
		return &Card{}, nil
	}

	currentCard := d.Deck[0]
	d.Deck = d.Deck[1:]
	return &currentCard, nil
}

func (d *Deck) Done() *Card {

}
