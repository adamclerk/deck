package deck

import "fmt"

// Card represents a Card with a Face and a Suit
type Card struct {
	face Face
	suit Suit
}

func (c Card) String() string {
	face := ""
	switch c.face {
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
	switch c.suit {
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
	return fmt.Sprintf("%x%x", c.face, c.suit)
}

// Compare compares 2 cards 1 if the passed in card is greater -1 if its lesser 0 of equal.
func (c *Card) Compare(k Card) int {
	if k.face > c.face {
		return 1
	}

	if k.face < c.face {
		return -1
	}

	if k.suit > c.suit {
		return 1
	}

	if k.suit < c.suit {
		return -1
	}

	return 0
}

//IsLessThan returns bool if card passed in is less then
func (c *Card) IsLessThan(k Card) bool {
	return c.Compare(k) == 1
}

//IsGreaterThan return bool if card passed in is greater then
func (c *Card) IsGreaterThan(k Card) bool {
	return c.Compare(k) == -1
}

//IsEqualTo returns true if the card is equal in face and
func (c *Card) IsEqualTo(k Card) bool {
	return c.Compare(k) == 0
}
