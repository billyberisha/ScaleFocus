package main

import (
	"fmt"
)

func main() {
	fmt.Println("Choose a card from 1 to 14, and a suit for that card")
	fmt.Println("Suits are as following: 4,3,2,1")
	compareCards(5, 4, 11, 4)
}

func compareCards(cOneVal int, cOneSuit int, cTwoVal int, cTwoSuit int) int {
	isValidCard := (cOneVal > 1 && cOneVal < 15) && (cTwoVal > 1 && cTwoVal < 15)
	isValidSuit := (cOneSuit > 0 && cOneSuit < 5) && (cTwoSuit > 0 && cTwoSuit < 5)
	if isValidCard && isValidSuit {
		if cOneVal > cTwoVal && cOneSuit > cTwoSuit {
			fmt.Println("Card one has won")
			return (1)
		} else if cOneVal < cTwoVal && cOneSuit < cTwoSuit {
			fmt.Println("Card two has won")
			return (-1)
		} else {
			fmt.Println("The game is even")
			return (0)
		}

	} else {
		fmt.Printf("The first card %v and the second card %v have to be in the range of 2 and 14\n", cOneVal, cTwoVal)
		fmt.Printf("The first suit %v and the second suit %v have to be in the range of 1 and 4\n", cOneSuit, cTwoSuit)

	}
	return 1
}
