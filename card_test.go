package deck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardToString(t *testing.T) {
	card := Card{ACE, HEART}
	result := fmt.Sprintf("%s", card)
	assert.Equal(t, "A♥", result, "These should be equal")
}

func TestCardSignature(t *testing.T) {
	card := Card{ACE, HEART}
	result := card.GetSignature()
	assert.Equal(t, "02", result, "These should be equal")
}

func TestCardCompare(t *testing.T) {
	card1 := Card{ACE, HEART}
	card2 := Card{KING, HEART}
	result := DefaultCompare(card1, card2)
	assert.Equal(t, -1, result, "These should be equal")
}

func TestCardIsLessThan(t *testing.T) {
	card1 := Card{ACE, HEART}
	card2 := Card{KING, HEART}
	result := DefaultCompare(card1, card2).IsLessThan()
	assert.Equal(t, true, result, "These should be equal")
}

func TestCardIsEqual(t *testing.T) {
	card1 := Card{ACE, HEART}
	card2 := Card{ACE, HEART}
	result := DefaultCompare(card1, card2).IsEqualTo()
	assert.Equal(t, true, result, "These should be equal")
}

func TestCardIsGreaterThan(t *testing.T) {
	card1 := Card{ACE, HEART}
	card2 := Card{KING, HEART}
	result := DefaultCompare(card2, card1).IsGreaterThan()
	assert.Equal(t, true, result, "These should be equal")
}

func TestCardIsGreaterThanWithSuit(t *testing.T) {
	card1 := Card{ACE, SPADE}
	card2 := Card{ACE, HEART}
	result := DefaultCompare(card1, card2).IsGreaterThan()
	assert.Equal(t, true, result, "These should be equal")
}

func TestCardIsLessThanWithSuit(t *testing.T) {
	card1 := Card{ACE, HEART}
	card2 := Card{ACE, SPADE}
	result := DefaultCompare(card1, card2).IsLessThan()
	assert.Equal(t, true, result, "These should be equal")
}
