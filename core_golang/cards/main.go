package main

//import "fmt"

//import "fmt"

//import "fmt"

func main() {
	//var card string = "Ace of spades"
	// card := "Ace of spades"
	// card = "Ace of hearts"
	// card = newCard()

	//slice
	// cards := deck{"Ace of spades", newCard()}
	// cards = append(cards, "queen of club")
	// fmt.Println(cards)
	// cards.printCards()

	//cards := newDeck()
	//cards.printCards()

	// hand,remainingDeck := deal(cards , 5)
	// hand.printCards()
	// remainingDeck.printCards()

	/* to get byte size
	greet := "Greeting!"
	fmt.Println([]byte(greet))
	*/

	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	/*
	cards := newDeckFromFile("my_cards")
	cards.printCards()
	*/
	cards := newDeck()
	cards.shuffle()
	cards.printCards()

}

func newCard() string {
	return "ace of diamonds"
}
