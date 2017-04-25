package deck

import "fmt"

func ExampleDebugf() {
	Debugf(true, "This will print")
	Debugf(false, "This will not print")
	//Output: This will print
}

func ExampleDecks() {
	deck, _ := New(Decks(2))

	fmt.Printf("Card Count: %d", deck.NumberOfCards())
	//Output: Card Count: 104
}

func ExampleEmpty() {
	deck, _ := New(Empty)

	fmt.Printf("Card Count: %d", deck.NumberOfCards())
	//Output: Card Count: 0
}

func ExampleFaces() {
	deck, _ := New(Faces(ACE), Unshuffled)

	fmt.Printf("Card Count: %d\n", deck.NumberOfCards())
	fmt.Printf("%s", deck)
	//OutPut:
	//Card Count: 4
	// A♣
	// A♦
	// A♥
	// A♠
}

func ExampleSuits() {
	deck, _ := New(Suits(HEART), Unshuffled)

	fmt.Printf("Card Count: %d\n", deck.NumberOfCards())
	fmt.Printf("%s", deck)
	//OutPut:
	//Card Count: 13
	// A♥
	// 2♥
	// 3♥
	// 4♥
	// 5♥
	// 6♥
	// 7♥
	// 8♥
	// 9♥
	// T♥
	// J♥
	// Q♥
	// K♥
}

// With unshuffled you always get the same card in a given position in the deck.
func ExampleUnshuffled() {
	deck, _ := New(Unshuffled)

	fmt.Printf("Card Count: %d\n", deck.NumberOfCards())
	fmt.Printf("First Card: %s\n", deck.Cards[0])
	fmt.Printf("Last Card: %s\n", deck.Cards[len(deck.Cards)-1])
	//Output:
	// Card Count: 52
	// First Card: A♣
	// Last Card: K♠
}

func ExampleWithCards() {
	deck, _ := New(WithCards(NewCard(ACE, HEART), NewCard(ACE, DIAMOND)), Unshuffled)

	fmt.Printf("Card Count: %d\n", deck.NumberOfCards())
	fmt.Printf("%s", deck)
	//Output:
	//Card Count: 2
	// A♥
	// A♦
}

func ExampleDeck_Deal() {
	deck, _ := New()
	player1Hand, _ := New(Empty)
	player2Hand, _ := New(Empty)

	//Let's play some poker: 5 card stud
	deck.Deal(5, player1Hand, player2Hand)

	fmt.Printf("Player1 Card Count: %d\n", player1Hand.NumberOfCards())
	fmt.Printf("Player2 Card Count: %d\n", player2Hand.NumberOfCards())
	fmt.Printf("Deck Card Count: %d\n", deck.NumberOfCards())
	//Output:
	// Player1 Card Count: 5
	// Player2 Card Count: 5
	// Deck Card Count: 42
}
