package main

import (
	cardgame "ScaleFocus/Lecture9/cardGame"
	carddraw "ScaleFocus/Lecture9/carddraw"
)

func main() {
	d := cardgame.NewDeck(55)
	slice1 := carddraw.DrawAllCards(d)

}
