package poker

// HandType represents whether a hand is a pair, suited, or offsuit
type HandType int

const (
    Pair HandType = iota
    Suited
    Offsuit
)

// CardCombination represents a specific combination of two cards with exact suits
type CardCombination struct {
    Card1 Card
    Card2 Card
    Selected bool  // This allows for partial selection within a hand type
}

// HandCombo represents a starting hand combination at the abstracted level (like AKs, QQ, etc.)
type HandCombo struct {
    Card1Value int
    Card2Value int
    Type       HandType
    Selected   bool
    
    // This will store all the specific card combinations this abstracted hand represents
    Combinations []CardCombination
}

// getCardName converts a card value to its name (helper function)
func getCardName(value int) string {
    for name, val := range Values {
        if val == value {
            return name
        }
    }
    return ""
}

// String returns the string representation (like "AKs", "TT", "76o")
func (h HandCombo) String() string {
    card1Name := getCardName(h.Card1Value)
    card2Name := getCardName(h.Card2Value)
    
    if h.Type == Pair {
        return card1Name + card2Name
    } else if h.Type == Suited {
        return card1Name + card2Name + "s"
    } else {
        return card1Name + card2Name + "o"
    }
}

// Range represents a collection of hand combinations
type Range struct {
    Grid      [13][13]HandCombo
    Combos    map[string]HandCombo // For quick lookups by string representation
    
    // For detailed analysis
    AllCombinations map[string]CardCombination // Key could be "As,Ks" format
    SelectedCombos  int // Track how many specific combinations are selected
}

// NewRange creates a new empty range
func NewRange() *Range {
    r := &Range{
        Combos: make(map[string]HandCombo),
        AllCombinations: make(map[string]CardCombination),
    }
    
    // Get card values in descending order (A to 2)
    values := []int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
    
    // Create a full deck to extract cards from
    deck := CreateDeck()
    
    // Fill the grid
    for i, v1 := range values {
        for j, v2 := range values {
            var combo HandCombo
            combo.Card1Value = v1
            combo.Card2Value = v2
            
            // Initialize specific combinations slice
            combo.Combinations = []CardCombination{}
            
            if i == j { // Pairs (diagonal)
                combo.Type = Pair
                
                // For pairs, find all 6 combinations (C(4,2) = 6 ways to choose 2 suits from 4)
                for i, suit1 := range Suits {
                    for _, suit2 := range Suits[i+1:]{
                        card1 := findCard(&deck, v1, suit1)
                        card2 := findCard(&deck, v1, suit2)
                        
                        cardCombo := CardCombination{
                            Card1: card1,
                            Card2: card2,
                            Selected: false,
                        }
                        
                        combo.Combinations = append(combo.Combinations, cardCombo)
                        comboKey := formatComboKey(card1, card2)
                        r.AllCombinations[comboKey] = cardCombo
                    }
                }
            } else if i < j { // Suited combos (upper triangle)
                combo.Type = Suited
                
                // For suited combos, there are 4 combinations (one per suit)
                for _, suit := range Suits {
                    card1 := findCard(&deck, v1, suit)
                    card2 := findCard(&deck, v2, suit)
                    
                    cardCombo := CardCombination{
                        Card1: card1,
                        Card2: card2,
                        Selected: false,
                    }
                    
                    combo.Combinations = append(combo.Combinations, cardCombo)
                    comboKey := formatComboKey(card1, card2)
                    r.AllCombinations[comboKey] = cardCombo
                }
            } else { // Offsuit combos (lower triangle)
                combo.Type = Offsuit
                
                // For offsuit combos, there are 12 combinations (4×3)
                for _, suit1 := range Suits {
                    for _, suit2 := range Suits {
                        if suit1 != suit2 {
                            card1 := findCard(&deck, v1, suit1)
                            card2 := findCard(&deck, v2, suit2)
                            
                            cardCombo := CardCombination{
                                Card1: card1,
                                Card2: card2,
                                Selected: false,
                            }
                            
                            combo.Combinations = append(combo.Combinations, cardCombo)
                            comboKey := formatComboKey(card1, card2)
                            r.AllCombinations[comboKey] = cardCombo
                        }
                    }
                }
            }
            
            r.Grid[i][j] = combo
            r.Combos[combo.String()] = combo
        }
    }
    
    return r
}

// Helper function to find a specific card in the deck
func findCard(deck *Deck, value int, suit string) Card {
    for _, card := range *deck {
        if card.Value == value && card.Suit == suit {
            return card
        }
    }
    return Card{} // Should never happen if deck is complete
}

// Format a consistent key for the combination map
func formatComboKey(card1, card2 Card) string {
    // Always put higher value card first
    if card1.Value < card2.Value {
        card1, card2 = card2, card1
    }
    return card1.Name + card1.Suit + "," + card2.Name + card2.Suit
}

func (r *Range) GetTotalCombinationsInRange() int {
    return r.SelectedCombos
}
