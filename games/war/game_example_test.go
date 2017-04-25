package war

import (
	"fmt"

	"github.com/adamclerk/deck"
)

// This example uses a stacked deck to get a specific result
func Example() {
	cards := []deck.Card{
		deck.NewCard(deck.ACE, deck.HEART),
		deck.NewCard(deck.ACE, deck.SPADE),
		deck.NewCard(deck.ACE, deck.HEART),
		deck.NewCard(deck.ACE, deck.SPADE),
		deck.NewCard(deck.TWO, deck.HEART),
		deck.NewCard(deck.ACE, deck.SPADE),
	}
	game, err := New(WithDeck(deck.Unshuffled, deck.WithCards(cards...)))
	if err != nil {
		panic(err)
	}
	err = game.Play()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Winner is: %s", game.Winner()[0].Name())
	// Output: Winner is: Player1
}

func ExampleCompare() {
	card1 := deck.NewCard(deck.TWO, deck.SPADE)
	card2 := deck.NewCard(deck.THREE, deck.HEART)

	result1 := Compare(card1, card2)
	result2 := Compare(card2, card1)
	result3 := Compare(card1, card1)

	fmt.Printf("Compare(%s, %s) Result is %d\n", card1, card2, result1)
	fmt.Printf("Compare(%s, %s) Result is %d\n", card2, card1, result2)
	fmt.Printf("Compare(%s, %s) Result is %d\n", card1, card1, result3)
	// Output:
	// Compare(2♠, 3♥) Result is -1
	// Compare(3♥, 2♠) Result is 1
	// Compare(2♠, 2♠) Result is 0
}

func ExampleDebug() {
	cards := []deck.Card{
		deck.NewCard(deck.ACE, deck.HEART),
		deck.NewCard(deck.ACE, deck.SPADE),
		deck.NewCard(deck.ACE, deck.HEART),
		deck.NewCard(deck.ACE, deck.SPADE),
		deck.NewCard(deck.TWO, deck.HEART),
		deck.NewCard(deck.ACE, deck.SPADE),
	}
	game, err := New(Debug, WithDeck(deck.Unshuffled, deck.WithCards(cards...)))
	if err != nil {
		panic(err)
	}
	err = game.Play()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Winner is: %s", game.Winner()[0].Name())
	// Output:
	// Game of War Started
	// PlayerHandEmptyEval P1: 3 P2: 3
	// Turn 1
	//   Card Count - P1: 3 P2: 3
	//   Cards on the Table
	//   P1: A♥  P2: A♠
	//   WAR!
	// PlayerHandEmptyEval P1: 1 P2: 1
	// Turn 2
	//   Card Count - P1: 1 P2: 1
	//   Cards on the Table
	//   P1: 2♥  P2: A♠
	//   Player 1 wins this round
	// PlayerHandEmptyEval P1: 6 P2: 0
	// Player hand empty
	// Winner is: Player1
}

func ExampleMaxTurns() {
	game, _ := New(MaxTurns(3))

	err := game.Play()
	fmt.Println(err)
	// Output: To many turns taken
}

func ExampleWithDeck() {
	cards := []deck.Card{
		deck.NewCard(deck.ACE, deck.HEART),
		deck.NewCard(deck.ACE, deck.SPADE),
		deck.NewCard(deck.ACE, deck.HEART),
		deck.NewCard(deck.ACE, deck.SPADE),
		deck.NewCard(deck.TWO, deck.HEART),
		deck.NewCard(deck.ACE, deck.SPADE),
	}

	game, _ := New(WithDeck(deck.WithCards(cards...), deck.Unshuffled))
	game.Play()
	fmt.Printf("Winner is: %s", game.Winner()[0].Name())
	// Output:
	// Winner is: Player1
}

func ExampleNew() {

}
