package deck

import "fmt"

// Deck is a deck of cards. An array of type Card
type Deck struct {
	cards []Card
}

func (d Deck) String() string {
	str := ""
	for _, card := range d.cards {
		str += fmt.Sprint(card) + "\n"
	}
	return str
}

// NewDeck creates and returns a new deck with the bool parameter to either shuffle (true) or non shuffle (false)
func NewDeck(shuffled bool) Deck {
	deck := NewSpecificDeck(shuffled, FACES, SUITS)
	return deck
}

// NewSpecificDeck creates and returns a deck that is created
// with all the premutations of an array of Faces and an array of Suits.
// The same bool parameter is expected to shuffle the deck
func NewSpecificDeck(shuffled bool, faces []Face, suits []Suit) Deck {
	cards := make([]Card, len(suits)*len(faces))
	for sindex, suit := range suits {
		for findex, face := range faces {
			index := (sindex * len(faces)) + findex
			cards[index] = Card{face, suit}
		}
	}
	deck := Deck{cards}
	if shuffled {
		deck.Shuffle()
	}
	return deck
}

// NewEmptyDeck creates an empty deck with an empty array of Cards
func NewEmptyDeck() Deck {
	deck := Deck{[]Card{}}
	return deck
}

// NumberOfCards is a utility function that tells you how many cards are left in the deck
func (d *Deck) NumberOfCards() int {
	return len(d.cards)
}

// Card is a getter function to retrive a specific card at a given index
func (d *Deck) Card(index int) *Card {
	return &d.cards[index]
}

// GetSignature returns the signature of the deck
// The signature is a string in which each card is
// represented as a hex character. Each hex character
// is in the same order as the deck
func (d *Deck) GetSignature() string {
	sig := ""
	for _, card := range d.cards {
		sig += card.GetSignature()
	}
	return sig
}
