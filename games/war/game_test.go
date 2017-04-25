package war

import (
	"math"
	"testing"

	"github.com/adamclerk/deck"
	"github.com/stretchr/testify/assert"
)

func TestWarGameStandStillFewCards(t *testing.T) {
	cards := []deck.Card{
		deck.Card{Face: deck.ACE, Suit: deck.HEART},
		deck.Card{Face: deck.ACE, Suit: deck.SPADE},
		deck.Card{Face: deck.ACE, Suit: deck.DIAMOND},
		deck.Card{Face: deck.ACE, Suit: deck.CLUB},
	}
	game, _ := New(
		WithDeck(
			deck.Unshuffled,
			deck.WithCards(cards...),
		),
		MaxTurns(4),
	)
	err := game.Play()
	assert.Equal(t, "Stalemate: No Cards left in players hands", err.Error())
}

func TestWarGameStandStill(t *testing.T) {
	cards := []deck.Card{
		deck.Card{Face: deck.ACE, Suit: deck.HEART},
		deck.Card{Face: deck.ACE, Suit: deck.SPADE},
		deck.Card{Face: deck.ACE, Suit: deck.DIAMOND},
		deck.Card{Face: deck.ACE, Suit: deck.CLUB},
		deck.Card{Face: deck.TWO, Suit: deck.DIAMOND},
		deck.Card{Face: deck.TWO, Suit: deck.CLUB},
	}
	game, _ := New(
		WithDeck(
			deck.Unshuffled,
			deck.WithCards(cards...),
		),
		MaxTurns(4),
	)
	err := game.Play()
	assert.Equal(t, "Stalemate: All Cards Match", err.Error())
}

func TestWarGameSimple(t *testing.T) {
	cards := []deck.Card{
		deck.Card{Face: deck.ACE, Suit: deck.HEART},
		deck.Card{Face: deck.TWO, Suit: deck.SPADE},
	}
	game, _ := New(
		WithDeck(
			deck.Unshuffled,
			deck.WithCards(cards...),
		),
		MaxTurns(4),
	)
	err := game.Play()
	assert.Equal(t, nil, err)
	assert.Equal(t, game.Winner()[0].Name(), "Player2")
}

func TestWarGameWithOneWar(t *testing.T) {
	cards := []deck.Card{
		deck.Card{Face: deck.ACE, Suit: deck.HEART},
		deck.Card{Face: deck.ACE, Suit: deck.SPADE},
		deck.Card{Face: deck.ACE, Suit: deck.HEART},
		deck.Card{Face: deck.ACE, Suit: deck.SPADE},
		deck.Card{Face: deck.TWO, Suit: deck.HEART},
		deck.Card{Face: deck.ACE, Suit: deck.SPADE},
	}
	game, _ := New(
		WithDeck(
			deck.Unshuffled,
			deck.WithCards(cards...),
		),
		MaxTurns(4),
	)
	err := game.Play()
	assert.Equal(t, nil, err)
	assert.Equal(t, game.Winner()[0].Name(), "Player1")
}

func BenchmarkWar(b *testing.B) {
	for n := 0; n < b.N; n++ {
		game, _ := New(
			MaxTurns(math.MaxInt64),
		)
		game.Play()
	}
}
