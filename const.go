package deck

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
