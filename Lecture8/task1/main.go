package main

import (
	"fmt"
	carddraw "golang/Lecture8/cardDraw"
	cardgame "golang/Lecture8/cardGame"
)

func main() {
	d := cardgame.NewDeck()
	slice := carddraw.DrawAllCards(d)
	a := 10
	fmt.Print(a)
	fmt.Println(slice)

}
