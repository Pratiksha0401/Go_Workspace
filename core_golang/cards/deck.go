package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of deck which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardsuits := []string{"Spades","Clubs", "Dimaonds","Hearts"}
	cardValues := []string{"Ace", "One","Six","Queen"}

	for _,suit := range cardsuits{
		for _,value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//receiver arg d is reference to copy of cards slice
func (d deck) printCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck , handSize int)(deck , deck){
	return d[:handSize],d[handSize:]
}

func (d deck) toString() string{
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}