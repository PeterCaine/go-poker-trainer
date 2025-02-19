package poker

import (
	"fmt"
	"testing"
)

func TestCreateDeck(t *testing.T){
    got := CreateDeck()
    if len(got) != 52 {
        t.Errorf("got %v != want: 52", got )
    }

// Check for unique cards
	cardSet := make(map[string]bool)
	for _, card := range got {
		cardKey := card.Name + card.Suit
		if cardSet[cardKey] {
			t.Errorf("Duplicate card found: %s of %s", card.Name, card.Suit)
		}
		cardSet[cardKey] = true
	}
}

func TestShuffleDeck(t *testing.T) {
	deck := CreateDeck()
    fmt.Print(len(deck))
	originalDeck := make([]Card, len(deck))
	copy(originalDeck, deck) // Copy original order before shuffling

	deck.ShuffleDeck()

	// Ensure deck length remains 52
	if len(deck) != 52 {
		t.Errorf("Deck length changed after shuffle. Expected 52, got %d", len(deck))
	}

	// Ensure all original cards are still present
	originalCards := make(map[string]bool)
	for _, card := range originalDeck {
		originalCards[card.Name+card.Suit] = true
	}

	for _, card := range deck {
		if !originalCards[card.Name+card.Suit] {
			t.Errorf("Card missing after shuffle: %s of %s", card.Name, card.Suit)
		}
	}

	// Ensure deck order changed (probabilistic)
	samePositionCount := 0
	for i := range deck {
		if deck[i] == originalDeck[i] {
			samePositionCount++
		}
	}

	if samePositionCount > 40 { // Allow some randomness, but expect significant change
		t.Errorf("Deck order did not change sufficiently after shuffle. %d cards remained in place", samePositionCount)
	}
}
