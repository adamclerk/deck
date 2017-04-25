package deck

import "fmt"

// Suit represents the suit of the card (spade, heart, diamon, club)
type Suit int

// Face represents the face of the card (ace, two...queen, king)
type Face int

// Contants for Suit ♠♥♦♣
const (
	CLUB Suit = iota
	DIAMOND
	HEART
	SPADE
)

// Contants for Face
const (
	ACE Face = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

// Global Variables representing the default suits and faces in a deck of cards
var (
	SUITS = []Suit{CLUB, DIAMOND, HEART, SPADE}
	FACES = []Face{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
)

// Card represents a playing card with a Face and a Suit
type Card struct {
	Face Face
	Suit Suit
}

func (c Card) String() string {
	face := ""
	switch c.Face {
	case 0:
		face = "A"
	case 1:
		face = "2"
	case 2:
		face = "3"
	case 3:
		face = "4"
	case 4:
		face = "5"
	case 5:
		face = "6"
	case 6:
		face = "7"
	case 7:
		face = "8"
	case 8:
		face = "9"
	case 9:
		face = "T"
	case 10:
		face = "J"
	case 11:
		face = "Q"
	case 12:
		face = "K"
	}
	suit := ""
	switch c.Suit {
	case 0:
		suit = "♣"
	case 1:
		suit = "♦"
	case 2:
		suit = "♥"
	case 3:
		suit = "♠"
	}
	return fmt.Sprintf("%s%s", face, suit)
}

// GetSignature is the hex representation of the Face and Suit of the card
func (c *Card) GetSignature() string {
	return fmt.Sprintf("%x%x", c.Face, c.Suit)
}
