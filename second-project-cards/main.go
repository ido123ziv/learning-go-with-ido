package main

import "fmt"

func main() {
	// classic way
	var cardClassic string = "Ace of Spades"
	fmt.Println(cardClassic)
	// go way
	gocard := "Ace of Spades" // note that this is applicable only on declaration
	gocard = "Five of Diamonds"
	fmt.Println(gocard)

	// from function

	cardFunction := newCard()
	fmt.Println(cardFunction)

	// slices
	cards := []string{"Ace of Clubs", newCard()}
	cards = append(cards, "Six of Spades") // append returns a new slice not modifing the current.

	fmt.Println(cards)

	// iterate over a slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// now we use the deck
	cardsDeck := deck{"Ace of Heart", newCard()}
	cardsDeck = append(cardsDeck, "six of Clubes")

	cardsDeck.print()

	fmt.Println("--------------------------")
	newCardsDeck := newDeck()
	newCardsDeck.print()

	hand, remainingCards := deal(newCardsDeck, 5)
	fmt.Println("Hand: ")
	hand.print()
	fmt.Println("Remaining: ")
	remainingCards.print()

}

func newCard() string {
	return "Three of Hearts"
}
