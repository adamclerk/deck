package deck

import (
	"math/rand"
	"time"
)

// MultiShuffle calls Shuffle multipule times
func (deck *Deck) MultiShuffle(iterations int) {
	for i := 0; i < iterations; i++ {
		deck.Shuffle()
	}
}

// Shuffle uses Knuth shuffle algo to randomize the deck in O(n) time
// sourced from https://gist.github.com/quux00/8258425
func (deck *Deck) Shuffle() {

	rand.Seed(time.Now().UnixNano())
	N := len(deck.cards)
	for i := 0; i < N; i++ {
		r := i + rand.Intn(N-i)
		deck.cards[r], deck.cards[i] = deck.cards[i], deck.cards[r]
	}
}
