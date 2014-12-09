package deck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallDeckToString(t *testing.T) {
	deck := NewEmptyDeck()
	deck.cards = append(deck.cards, Card{ACE, HEART}, Card{KING, HEART})
	result := fmt.Sprintf("%s", deck)
	assert.Equal(t, "A♥\nK♥\n", result, "These should be equal")
}

func TestDeckToString(t *testing.T) {
	deck := NewDeck(false)
	result := fmt.Sprintf("%s", deck)
	assert.Equal(t, "A♣\n2♣\n3♣\n4♣\n5♣\n6♣\n7♣\n8♣\n9♣\nT♣\nJ♣\nQ♣\nK♣\n"+
		"A♦\n2♦\n3♦\n4♦\n5♦\n6♦\n7♦\n8♦\n9♦\nT♦\nJ♦\nQ♦\nK♦\n"+
		"A♥\n2♥\n3♥\n4♥\n5♥\n6♥\n7♥\n8♥\n9♥\nT♥\nJ♥\nQ♥\nK♥\n"+
		"A♠\n2♠\n3♠\n4♠\n5♠\n6♠\n7♠\n8♠\n9♠\nT♠\nJ♠\nQ♠\nK♠\n", result, "These should be equal")
}

func TestDeckSignature(t *testing.T) {
	deck := NewEmptyDeck()
	deck.cards = append(deck.cards, Card{ACE, HEART}, Card{KING, HEART}, Card{TEN, CLUB})
	result := deck.GetSignature()
	assert.Equal(t, "02c290", result, "These should be equal")
}

func TestCompleteDeckSignature(t *testing.T) {
	deck := NewDeck(false)
	result := deck.GetSignature()
	assert.Equal(t, "00102030405060708090a0b0c001112131415161718191a1b1c102122232425262728292a2b2c203132333435363738393a3b3c3", result, "These should be equal")
}

func TestNewTinyDeck(t *testing.T) {
	deck := NewSpecificDeck(true, FACES, []Suit{SPADE})
	assert.Equal(t, 13, deck.NumberOfCards())
}

func TestNewSmallDeck(t *testing.T) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE, HEART})
	assert.Equal(t, 26, deck.NumberOfCards())
}

func BenchmarkTinyDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE})
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkSmallDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE, HEART})
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkMediumDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, []Suit{SPADE, HEART, DIAMOND})
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkDeckShuffle(b *testing.B) {
	deck := NewSpecificDeck(false, FACES, SUITS)
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}
