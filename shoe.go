package deck

import "fmt"

// Shoe is a collection of Decks. Used in games like blackjack
type Shoe struct {
	decks []Deck
}

func (s Shoe) String() string {
	str := ""
	if len(s.decks) > 0 {
		for i := 0; i < len(s.decks); i++ {
			str += "Deck-"
			str += fmt.Sprint(i + 1)
			str += "\t"
		}
		str += "\n"

		for c := 0; c < len(s.decks[0].cards); c++ {
			for d := 0; d < len(s.decks); d++ {
				str += fmt.Sprint(s.Deck(d).Card(c)) + "\t"
			}
			str += "\n"
		}
	}
	return str
}

// NewShoe creates and returns a new shoe of decks
func NewShoe(shuffled bool, decks int) Shoe {
	shoe := Shoe{[]Deck{}}
	for i := 0; i < decks; i++ {
		shoe.decks = append(shoe.decks, NewDeck(shuffled))
	}
	return shoe
}

// NewSpecificShoe creates and returns a new show of decks based on custom suits and faces
func NewSpecificShoe(shuffled bool, decks int, faces []Face, suits []Suit) Shoe {
	shoe := Shoe{[]Deck{}}
	for i := 0; i < decks; i++ {
		shoe.decks = append(shoe.decks, NewSpecificDeck(shuffled, faces, suits))
	}

	return shoe
}

// NewEmptyShoe returns an empty shoe
func NewEmptyShoe() Shoe {
	shoe := Shoe{}
	return shoe
}

// Deck is a getter for the private deck
func (s *Shoe) Deck(index int) *Deck {
	return &s.decks[index]
}

// NumberOfDecks is a utility function to get the total number of decks in the shoe
func (s *Shoe) NumberOfDecks() int {
	return len(s.decks)
}
