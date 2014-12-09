package deck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	expected := NewDeck(false)
	deck := NewDeck(false)
	deck.Shuffle()
	assert.NotEqual(t, expected.GetSignature(), deck.GetSignature())
}

func TestMultiShuffle(t *testing.T) {
	expected := NewDeck(false)
	deck := NewDeck(false)
	deck.MultiShuffle(10)
	assert.NotEqual(t, expected.GetSignature(), deck.GetSignature())
}
