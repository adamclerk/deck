package war

import (
	"errors"

	"github.com/adamclerk/deck"
)

// War rules can be found here: http://www.bicyclecards.com/how-to-play/war/
type War struct {
	debug              bool
	turns              int
	maxTurns           int
	player1, player2   Player
	deck, pile1, pile2 *deck.Deck
	winner             *Player
	compare            func(i, j deck.Card) deck.CompareResult
}

// Options how to configure a game of war
type Options struct {
	DeckOptions   []func(*deck.Options)
	MaxTurns      int
	Debug         bool
	CustomCompare func(i, j deck.Card) deck.CompareResult
}

// Play plays the game
func (w *War) Play() error {
	deck.Debugf(w.debug, "Game of War Started\n")
	// The Deal
	// --------
	// The deck is divided evenly, with each player receiving 26 cards, dealt one at a time, face down.
	// Anyone may deal first. Each player places his stack of cards face down, in front of him
	w.deck.Deal(w.deck.NumberOfCards()/2, &w.player1.hand, &w.player2.hand)

	// The Play
	// --------
	// Each player turns up a card at the same time and the player with the higher card takes both cards and puts them, face down, on the bottom of his stack.
	// If the cards are the same rank, it is War. Each player turns up one card face down and one card face up.
	// The player with the higher cards takes both piles (six cards).
	// If the turned-up cards are again the same rank, each player places another card face down and turns another card face up.
	// The player with the higher card takes all 10 cards, and so on.
	for !w.PlayerHandEmpty() {
		w.turns++
		if w.turns > w.maxTurns {
			return errors.New("To many turns taken")
		}

		if w.player1.hand.NumberOfCards() == 0 && w.player1.hand.NumberOfCards() == 0 { //stalemate
			return errors.New("Stalemate: All Cards Match")
		}

		deck.Debugf(w.debug, "Turn %d\n", w.turns)
		deck.Debugf(w.debug, "  Card Count - P1: %d P2: %d\n", w.player1.hand.NumberOfCards(), w.player2.hand.NumberOfCards())
		// Each player turns up a card at the same time
		w.player1.hand.Deal(1, w.pile1)
		w.player2.hand.Deal(1, w.pile2)
		// the player with the higher card takes both cards and puts them, face down, on the bottom of his stack
		p1Card := w.pile1.Cards[len(w.pile1.Cards)-1]
		p2Card := w.pile2.Cards[len(w.pile2.Cards)-1]

		deck.Debugf(w.debug, "  Cards on the Table\n")
		deck.Debugf(w.debug, "  P1: %s  P2: %s\n", p1Card, p2Card)

		if w.compare(p1Card, p2Card).IsGreaterThan() {
			deck.Debugf(w.debug, "  Player 1 wins this round\n")
			w.pile1.Deal(len(w.pile1.Cards), &w.player1.hand)
			w.pile2.Deal(len(w.pile2.Cards), &w.player1.hand)
		} else if w.compare(p2Card, p1Card).IsGreaterThan() {
			deck.Debugf(w.debug, "  Player 2 wins this round\n")
			w.pile1.Deal(len(w.pile1.Cards), &w.player2.hand)
			w.pile2.Deal(len(w.pile2.Cards), &w.player2.hand)
		} else { // If the cards are the same rank, it is War.
			deck.Debugf(w.debug, "  WAR!\n")
			if w.player1.hand.NumberOfCards() == 0 && w.player1.hand.NumberOfCards() == 0 { //stalemate
				return errors.New("Stalemate: All Cards Match")
			}
			w.player1.hand.Deal(1, w.pile1)
			w.player2.hand.Deal(1, w.pile2)

			if w.player1.hand.NumberOfCards() == 0 && w.player1.hand.NumberOfCards() == 0 { //stalemate
				return errors.New("Stalemate: No Cards left in players hands")
			}
		}
	}

	deck.Debugf(w.debug, "Player hand empty\n")

	// How to Keep Score
	// -----------------
	// The game ends when one player has won all the cards
	if w.player1.hand.NumberOfCards() > 0 {
		w.winner = &w.player1
		return nil
	} else if w.player2.hand.NumberOfCards() > 0 {
		w.winner = &w.player2
		return nil
	} else {
		return errors.New("No winner found")
	}
}

// PlayerHandEmpty checks to see if one of the players is out of cards
func (w *War) PlayerHandEmpty() bool {
	deck.Debugf(w.debug, "PlayerHandEmptyEval P1: %d P2: %d\n", w.player1.hand.NumberOfCards(), w.player2.hand.NumberOfCards())
	return w.player1.hand.NumberOfCards() == 0 || w.player2.hand.NumberOfCards() == 0
}

// Winner announces the winner
func (w *War) Winner() []Player {
	return []Player{*w.winner}
}

// Player is the player for the game of War
type Player struct {
	hand       deck.Deck
	playerName string
}

// Name returns the players name for verification and announcement
func (p Player) Name() string {
	return p.playerName
}

// WithDeck allows the a game to be configured with a specific deck
func WithDeck(deckOptions ...func(*deck.Options)) func(*Options) {
	return func(o *Options) {
		o.DeckOptions = deckOptions
	}
}

// MaxTurns before returning an error.
// There could be decks that are ordered in a such a way as to create never ending games of WAR.
// This is to put a limit on the amount of time spent playing.
func MaxTurns(turns int) func(*Options) {
	return func(o *Options) {
		o.MaxTurns = turns
	}
}

// Debug sets the debug param for the game
func Debug(o *Options) {
	o.Debug = true
}

// Compare compares 2 cards.
// returns 1 if i is greater.
// returns  -1 if it's lesser.
// returns 0 of equal.
func Compare(i, j deck.Card) deck.CompareResult {
	if i.Face() > j.Face() {
		return 1
	}

	if i.Face() < j.Face() {
		return -1
	}

	return 0
}

// New function creates a new game of War
func New(options ...func(*Options)) (*War, error) {
	opt := Options{DeckOptions: []func(*deck.Options){}, MaxTurns: 1000000, Debug: false, CustomCompare: Compare}

	for _, option := range options {
		option(&opt)
	}

	d, _ := deck.New(opt.DeckOptions...)
	p1, _ := deck.New(deck.Empty)
	p2, _ := deck.New(deck.Empty)

	return &War{
		compare:  opt.CustomCompare,
		debug:    opt.Debug,
		maxTurns: opt.MaxTurns,
		turns:    0,
		deck:     d,
		player1:  Player{playerName: "Player1"},
		player2:  Player{playerName: "Player2"},
		pile1:    p1,
		pile2:    p2,
	}, nil
}
