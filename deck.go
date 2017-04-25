// Package deck is a library that describe the card game domain
package deck

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Deck is a deck of cards. An array of type Card
type Deck struct {
	Cards         []Card
	NumberOfDecks int
}

// Options is the struct used to describe now a Deck should be created
type Options struct {
	Shuffled  bool   // should the deck be shuffled
	Cards     []Card // default array of cards to use
	Faces     []Face
	Suits     []Suit
	Decks     int
	Signature string
}

// New creates a new deck based on Options
func New(options ...func(*Options)) (*Deck, error) {
	opt := Options{Shuffled: true, Faces: FACES, Suits: SUITS, Decks: 1, Cards: []Card{}, Signature: ""}
	for _, option := range options {
		option(&opt)
	}

	cards := opt.Cards

	if opt.Signature != "" {
		convertSignature(opt.Signature, &cards)
	}

	if len(cards) == 0 {
		cards = make([]Card, len(opt.Suits)*len(opt.Faces)*opt.Decks)
		index := 0
		for i := 0; i < opt.Decks; i++ {
			for _, suit := range opt.Suits {
				for _, face := range opt.Faces {
					cards[index] = NewCard(face, suit)
					index++
				}
			}
		}
	}
	deck := Deck{cards, opt.Decks}
	if opt.Shuffled {
		deck.Shuffle()
	}
	return &deck, nil
}

// Seed is used to seed rng. It should be called only once before creating any deck.
// A single invocation is good enough on each executable invocation.
func Seed() {
	rand.Seed(time.Now().UnixNano())
}

// WithCards creates a deck with specific cards.
// Paired with he Unshuffled option and you can stack a deck for testing.
// With cards takes precedence over Suits and Faces Options.
func WithCards(cards ...Card) func(*Options) {
	return func(o *Options) {
		o.Cards = cards
	}
}

// Decks is a functional option used to create a shoe with multiple decks.
func Decks(count int) func(*Options) {
	return func(o *Options) {
		o.Decks = count
	}
}

// Empty is a functional option used to create an empty deck.
func Empty(o *Options) {
	o.Faces = []Face{}
	o.Suits = []Suit{}
}

// Unshuffled is a functional option used to stop the default shuffle used when all decks are created.
// This is mostly used for testing.
func Unshuffled(o *Options) {
	o.Shuffled = false
}

// Suits is a functional option used to provide the suits that should be included in a created deck.
func Suits(suits ...Suit) func(*Options) {
	return func(o *Options) {
		o.Suits = suits
	}
}

// Faces is a functional option used to provide the faces that should be included in a created deck.
func Faces(faces ...Face) func(*Options) {
	return func(o *Options) {
		o.Faces = faces
	}
}

// FromSignature is a functional option used to create decks from a given hex signature
func FromSignature(sig string) func(*Options) {
	return func(o *Options) {
		o.Signature = sig
	}
}

func (d *Deck) String() string {
	str := ""
	for _, card := range d.Cards {
		str += fmt.Sprint(card) + "\n"
	}
	return str
}

// NumberOfCards is a utility function that tells you how many cards are left in the deck
func (d *Deck) NumberOfCards() int {
	return len(d.Cards)
}

// Deal distributes cards to other decks/hands
func (d *Deck) Deal(cards int, hands ...*Deck) {
	for i := 0; i < cards; i++ {
		for _, hand := range hands {
			card := d.Cards[0]
			d.Cards = d.Cards[1:]
			hand.Cards = append(hand.Cards, card)
		}
	}
}

// Shuffle uses Knuth shuffle algo to randomize the deck in O(n) time
// sourced from https://gist.github.com/quux00/8258425
func (d *Deck) Shuffle() {
	N := len(d.Cards)
	for i := 0; i < N; i++ {
		r := i + rand.Intn(N-i)
		d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
	}
}

// ShufflePerm uses rand.Perm instead of the many calls to rand.Intn.
//  When compared to the current implementation:
//
//  benchmark                        old ns/op     new ns/op     delta
//  BenchmarkTinyDeckShuffle-8       524           537           +2.48%
//  BenchmarkSmallDeckShuffle-8      1119          1070          -4.38%
//  BenchmarkMediumDeckShuffle-8     1611          1626          +0.93%
//  BenchmarkDeckShuffle-8           2115          2194          +3.74%
//  BenchmarkLargeDeckShuffle-8      21301         21408         +0.50%
//
//  Conclusion: Not Recommended
func (d *Deck) ShufflePerm() {
	N := len(d.Cards)
	perm := rand.Perm(N)
	for i := 0; i < N; i++ {
		d.Cards[perm[i]], d.Cards[i] = d.Cards[i], d.Cards[perm[i]]
	}
}

// GetSignature returns the signature of the deck
// The signature is a string in which each card is
// represented as a hex character. Each hex character
// is in the same order as the deck
func (d *Deck) GetSignature() string {
	sig := ""
	for _, card := range d.Cards {

		sig += card.GetSignature()
	}
	return sig
}

func convertSignature(sig string, cards *[]Card) {
	for i := 0; i < len(sig)-1; i = i + 2 {
		face, _ := strconv.ParseInt(string(sig[i]), 16, 8)
		suit, _ := strconv.ParseInt(string(sig[i+1]), 16, 8)
		*cards = append(*cards, NewCard(Face(face), Suit(suit)))
	}
	// return cards
}

// CompareResult is the custom type returned when comparing cards
type CompareResult int

// DefaultCompare is the default comparison function
// Currently not used in any games.
func DefaultCompare(i, j Card) CompareResult {
	if i.Face() > j.Face() {
		return 1
	}

	if i.Face() < j.Face() {
		return -1
	}

	if i.Suit() > j.Suit() {
		return 1
	}

	if i.Suit() < j.Suit() {
		return -1
	}

	return 0
}

// IsGreaterThan is a utility function to make code more readable when comparing cards.
// This is used in conjunction with DefaultCompare or any comparison function that returns CompareResult
func (i CompareResult) IsGreaterThan() bool {
	return i == 1
}

// IsLessThan is a utility function to make code more readable when comparing cards.
// This is used in conjunction with DefaultCompare or any comparison function that returns CompareResult
func (i CompareResult) IsLessThan() bool {
	return i == -1
}

// IsEqualTo is a utility function to make code more readable when comparing cards.
// This is used in conjunction with DefaultCompare or any comparison function that returns CompareResult
func (i CompareResult) IsEqualTo() bool {
	return i == 0
}
