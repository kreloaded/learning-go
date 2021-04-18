package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

const DEFAULT_FILE_PERMISSION = 0666

type deck []string

/**
 * Create a new deck.
 */
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/**
 * Print deck contents.
 */
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/**
 * Get a deal from the deck.
 */
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/**
 * Convert deck to string.
 */
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

/**
 * Save deck to file in local memory.
 */
func (d deck) saveToFile(fileName string) error {
	byteSlice := []byte(d.toString())

	return ioutil.WriteFile(fileName, byteSlice, DEFAULT_FILE_PERMISSION)
}

/**
 * Read a new deck from file(local memory).
 */
func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), ",")

	return deck(s)
}

/**
 * Shuffle deck using random numbers.
 */
func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

	return d
}
