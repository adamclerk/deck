package deck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShoeToString(t *testing.T) {
	shoe := NewShoe(false, 2)
	result := fmt.Sprintf("%s", shoe)
	assert.Equal(t, "Deck-1\tDeck-2\t\nA♣\tA♣\t\n2♣\t2♣\t\n3♣\t3♣\t\n4♣\t4♣\t\n5♣\t5♣\t\n"+
		"6♣\t6♣\t\n7♣\t7♣\t\n8♣\t8♣\t\n9♣\t9♣\t\nT♣\tT♣\t\nJ♣\tJ♣\t\nQ♣\tQ♣\t\nK♣\tK♣\t\n"+
		"A♦\tA♦\t\n2♦\t2♦\t\n3♦\t3♦\t\n4♦\t4♦\t\n5♦\t5♦\t\n6♦\t6♦\t\n7♦\t7♦\t\n8♦\t8♦\t\n"+
		"9♦\t9♦\t\nT♦\tT♦\t\nJ♦\tJ♦\t\nQ♦\tQ♦\t\nK♦\tK♦\t\nA♥\tA♥\t\n2♥\t2♥\t\n3♥\t3♥\t\n"+
		"4♥\t4♥\t\n5♥\t5♥\t\n6♥\t6♥\t\n7♥\t7♥\t\n8♥\t8♥\t\n9♥\t9♥\t\nT♥\tT♥\t\nJ♥\tJ♥\t\n"+
		"Q♥\tQ♥\t\nK♥\tK♥\t\nA♠\tA♠\t\n2♠\t2♠\t\n3♠\t3♠\t\n4♠\t4♠\t\n5♠\t5♠\t\n6♠\t6♠\t\n"+
		"7♠\t7♠\t\n8♠\t8♠\t\n9♠\t9♠\t\nT♠\tT♠\t\nJ♠\tJ♠\t\nQ♠\tQ♠\t\nK♠\tK♠\t\n", result, "These should be equal")
}

func TestNewSpecificShoe(t *testing.T) {
	shoe := NewSpecificShoe(false, 2, []Face{ACE, TWO}, []Suit{HEART, SPADE})
	result := fmt.Sprintf("%s", shoe)
	assert.Equal(t, "Deck-1\tDeck-2\t\nA♥\tA♥\t\n2♥\t2♥\t\nA♠\tA♠\t\n2♠\t2♠\t\n", result, "These should be equal")
}

func TestEmptyShoe(t *testing.T) {
	shoe := NewEmptyShoe()
	result := shoe.NumberOfDecks()
	assert.Equal(t, 0, result, "These should be equal")
}
