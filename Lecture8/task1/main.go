package main

import (
	carddraw "ScaleFocus/Lecture8/cardDraw"
	cardgame "ScaleFocus/Lecture8/cardGame"
	"fmt"
)

func main() {
	d := cardgame.NewDeck()
	slice := carddraw.DrawAllCards(d)
	a := 10
	fmt.Print(a)
	fmt.Println(slice)

}
