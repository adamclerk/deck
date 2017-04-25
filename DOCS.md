# deck
--
    import "github.com/adamclerk/deck"

Package deck is a library that describe the card game domain

## Usage

```go
var (
	SUITS = []Suit{CLUB, DIAMOND, HEART, SPADE}
	FACES = []Face{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
)
```
Global Variables representing the default suits and faces in a deck of cards

#### func  Debugf

```go
func Debugf(debug bool, format string, a ...interface{}) (int, error)
```
Debugf statement for use to debug games

#### func  Decks

```go
func Decks(count int) func(*Options)
```
Decks is a functional option used to create a shoe with multiple decks.

#### func  Empty

```go
func Empty(o *Options)
```
Empty is a functional option used to create an empty deck.

#### func  Faces

```go
func Faces(faces ...Face) func(*Options)
```
Faces is a functional option used to provide the faces that should be included
in a created deck.

#### func  FromSignature

```go
func FromSignature(sig string) func(*Options)
```
FromSignature is a functional option used to create decks from a given hex
signature

#### func  Seed

```go
func Seed()
```
Seed is used to seed rng. It should be called only once before creating any
deck. A single invocation is good enough on each executable invocation.

#### func  Suits

```go
func Suits(suits ...Suit) func(*Options)
```
Suits is a functional option used to provide the suits that should be included
in a created deck.

#### func  Unshuffled

```go
func Unshuffled(o *Options)
```
Unshuffled is a functional option used to stop the default shuffle used when all
decks are created. This is mostly used for testing.

#### func  WithCards

```go
func WithCards(cards ...Card) func(*Options)
```
WithCards creates a deck with specific cards. Paired with he Unshuffled option
and you can stack a deck for testing. With cards takes precedence over Suits and
Faces Options.

#### type Card

```go
type Card struct {
	Face Face
	Suit Suit
}
```

Card represents a playing card with a Face and a Suit

#### func (*Card) GetSignature

```go
func (c *Card) GetSignature() string
```
GetSignature is the hex representation of the Face and Suit of the card

#### func (Card) String

```go
func (c Card) String() string
```

#### type CompareResult

```go
type CompareResult int
```

CompareResult is the custom type returned when comparing cards

#### func  DefaultCompare

```go
func DefaultCompare(i, j Card) CompareResult
```
DefaultCompare is the default comparison function Currently not used in any
games.

#### func (CompareResult) IsEqualTo

```go
func (i CompareResult) IsEqualTo() bool
```
IsEqualTo is a utility function to make code more readable when comparing cards.
This is used in conjunction with DefaultCompare or any comparison function that
returns CompareResult

#### func (CompareResult) IsGreaterThan

```go
func (i CompareResult) IsGreaterThan() bool
```
IsGreaterThan is a utility function to make code more readable when comparing
cards. This is used in conjunction with DefaultCompare or any comparison
function that returns CompareResult

#### func (CompareResult) IsLessThan

```go
func (i CompareResult) IsLessThan() bool
```
IsLessThan is a utility function to make code more readable when comparing
cards. This is used in conjunction with DefaultCompare or any comparison
function that returns CompareResult

#### type Deck

```go
type Deck struct {
	Cards         []Card
	NumberOfDecks int
}
```

Deck is a deck of cards. An array of type Card

#### func  New

```go
func New(options ...func(*Options)) (*Deck, error)
```
New creates a new deck based on Options

#### func (*Deck) Deal

```go
func (d *Deck) Deal(cards int, hands ...*Deck)
```
Deal distributes cards to other decks/hands

#### func (*Deck) GetSignature

```go
func (d *Deck) GetSignature() string
```
GetSignature returns the signature of the deck The signature is a string in
which each card is represented as a hex character. Each hex character is in the
same order as the deck

#### func (*Deck) NumberOfCards

```go
func (d *Deck) NumberOfCards() int
```
NumberOfCards is a utility function that tells you how many cards are left in
the deck

#### func (*Deck) Shuffle

```go
func (d *Deck) Shuffle()
```
Shuffle uses Knuth shuffle algo to randomize the deck in O(n) time sourced from
https://gist.github.com/quux00/8258425

#### func (*Deck) ShufflePerm

```go
func (d *Deck) ShufflePerm()
```
ShufflePerm uses rand.Perm instead of the many calls to rand.Intn.

    When compared to the current implementation:

    benchmark                        old ns/op     new ns/op     delta
    BenchmarkTinyDeckShuffle-8       524           537           +2.48%
    BenchmarkSmallDeckShuffle-8      1119          1070          -4.38%
    BenchmarkMediumDeckShuffle-8     1611          1626          +0.93%
    BenchmarkDeckShuffle-8           2115          2194          +3.74%
    BenchmarkLargeDeckShuffle-8      21301         21408         +0.50%

    Conclusion: Not Recommended

#### func (*Deck) String

```go
func (d *Deck) String() string
```

#### type Face

```go
type Face int
```

Face represents the face of the card (ace, two...queen, king)

```go
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
```
Constants for Face

#### type Options

```go
type Options struct {
	Shuffled  bool   // should the deck be shuffled
	Cards     []Card // default array of cards to use
	Faces     []Face
	Suits     []Suit
	Decks     int
	Signature string
}
```

Options is the struct used to describe now a Deck should be created

#### type Suit

```go
type Suit int
```

Suit represents the suit of the card (spade, heart, diamond, club)

```go
const (
	CLUB Suit = iota
	DIAMOND
	HEART
	SPADE
)
```
Constants for Suit ♠♥♦♣
