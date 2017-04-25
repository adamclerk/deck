package deck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallDeckToString(t *testing.T) {
	deck, _ := New(Empty)
	deck.Cards = append(deck.Cards, Card{ACE, HEART}, Card{KING, HEART})
	result := fmt.Sprintf("%s", deck)
	assert.Equal(t, "A♥\nK♥\n", result, "These should be equal")
}

func TestDeckToString(t *testing.T) {
	deck, _ := New(Unshuffled)
	result := fmt.Sprintf("%s", deck)
	assert.Equal(t, "A♣\n2♣\n3♣\n4♣\n5♣\n6♣\n7♣\n8♣\n9♣\nT♣\nJ♣\nQ♣\nK♣\n"+
		"A♦\n2♦\n3♦\n4♦\n5♦\n6♦\n7♦\n8♦\n9♦\nT♦\nJ♦\nQ♦\nK♦\n"+
		"A♥\n2♥\n3♥\n4♥\n5♥\n6♥\n7♥\n8♥\n9♥\nT♥\nJ♥\nQ♥\nK♥\n"+
		"A♠\n2♠\n3♠\n4♠\n5♠\n6♠\n7♠\n8♠\n9♠\nT♠\nJ♠\nQ♠\nK♠\n", result, "These should be equal")
}

func TestDeckSignature(t *testing.T) {
	deck, _ := New(Empty)
	deck.Cards = append(deck.Cards, Card{ACE, HEART}, Card{KING, HEART}, Card{TEN, CLUB})
	result := deck.GetSignature()
	assert.Equal(t, "02c290", result, "These should be equal")
}

func TestCompleteDeckSignature(t *testing.T) {
	deck, _ := New(Unshuffled)
	result := deck.GetSignature()
	assert.Equal(t, "00102030405060708090a0b0c001112131415161718191a1b1c102122232425262728292a2b2c203132333435363738393a3b3c3", result, "These should be equal")
}

func TestNewTinyDeck(t *testing.T) {
	deck, _ := New(Unshuffled, Suits(SPADE))
	assert.Equal(t, 13, deck.NumberOfCards())
}

func TestNewSmallDeck(t *testing.T) {
	deck, _ := New(Unshuffled, Suits(SPADE, HEART))
	assert.Equal(t, 26, deck.NumberOfCards())
}

func TestDeal(t *testing.T) {
	deck, _ := New(Unshuffled, Suits(SPADE), Faces(FACES...))
	assert.Equal(t, deck.NumberOfCards(), 13)
	hand1, _ := New(Empty)
	hand2, _ := New(Empty)
	deck.Deal(5, hand1, hand2)
	assert.Equal(t, deck.NumberOfCards(), 3)
	assert.Equal(t, hand1.NumberOfCards(), 5)
	assert.Equal(t, hand2.NumberOfCards(), 5)
	assert.Equal(t, hand1.String(), "A♠\n3♠\n5♠\n7♠\n9♠\n")
	assert.Equal(t, hand2.String(), "2♠\n4♠\n6♠\n8♠\nT♠\n")
}

func TestShuffle(t *testing.T) {
	expected, _ := New(Unshuffled)
	deck, _ := New(Unshuffled)
	deck.Shuffle()
	assert.NotEqual(t, expected.GetSignature(), deck.GetSignature())
}

func TestShufflePerm(t *testing.T) {
	expected, _ := New(Unshuffled)
	deck, _ := New(Unshuffled)
	deck.ShufflePerm()
	assert.NotEqual(t, expected.GetSignature(), deck.GetSignature())
}

func TestWithCards(t *testing.T) {
	cards := []Card{
		Card{ACE, HEART},
		Card{ACE, SPADE},
		Card{ACE, DIAMOND},
		Card{ACE, CLUB},
	}
	deck, _ := New(WithCards(cards...), Unshuffled)
	assert.Equal(t, deck.NumberOfCards(), 4)
	assert.Equal(t, deck.Cards[0].String(), "A♥")
	assert.Equal(t, deck.Cards[1].String(), "A♠")
	assert.Equal(t, deck.Cards[2].String(), "A♦")
	assert.Equal(t, deck.Cards[3].String(), "A♣")
}

func BenchmarkTinyDeckShuffle(b *testing.B) {
	deck, _ := New(Unshuffled, Suits(SPADE))
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkSmallDeckShuffle(b *testing.B) {
	deck, _ := New(Unshuffled, Faces(FACES...), Suits(SPADE, HEART))
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkMediumDeckShuffle(b *testing.B) {
	deck, _ := New(Unshuffled, Faces(FACES...), Suits(SPADE, HEART, DIAMOND))
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}

func BenchmarkDeckShuffle(b *testing.B) {
	deck, _ := New(Unshuffled, Faces(FACES...), Suits(SUITS...))
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}
func BenchmarkLargeDeckShuffle(b *testing.B) {
	deck, _ := New(Unshuffled, Decks(10), Faces(FACES...), Suits(SUITS...))
	for n := 0; n < b.N; n++ {
		deck.Shuffle()
	}
}
func BenchmarkLargeDeckShuffleAlt(b *testing.B) {
	deck, _ := New(Unshuffled, Decks(10), Faces(FACES...), Suits(SUITS...))
	for n := 0; n < b.N; n++ {
		deck.ShufflePerm()
	}
}
func TestShoeToString(t *testing.T) {
	shoe, _ := New(Unshuffled, Decks(1))
	result := fmt.Sprintf("%s", shoe)
	assert.Equal(t, "A♣\n2♣\n3♣\n4♣\n5♣\n6♣\n7♣\n8♣\n9♣\nT♣\nJ♣\nQ♣\nK♣\nA♦\n2♦\n3♦\n4♦\n5♦\n6♦\n7♦\n8♦\n9♦\nT♦\nJ♦\nQ♦\nK♦\nA♥\n2♥\n3♥\n4♥\n5♥\n6♥\n7♥\n8♥\n9♥\nT♥\nJ♥\nQ♥\nK♥\nA♠\n2♠\n3♠\n4♠\n5♠\n6♠\n7♠\n8♠\n9♠\nT♠\nJ♠\nQ♠\nK♠\n", result, "These should be equal")
}

func TestNewSpecificShoe(t *testing.T) {
	shoe, _ := New(Unshuffled, Decks(2), Faces(ACE), Suits(HEART))
	result := fmt.Sprintf("%s", shoe)
	assert.Equal(t, "A♥\nA♥\n", result, "These should be equal")
}

func TestEmptyShoe(t *testing.T) {
	shoe, _ := New(Decks(0))
	result := shoe.NumberOfDecks
	assert.Equal(t, 0, result, "These should be equal")
}
