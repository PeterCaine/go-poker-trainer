package poker

import "math/rand"


type Card struct {
    Value   int
    Name    string
    Suit    string
}


var Suits = []string {
    "heart",
    "diamond", 
    "spade", 
    "club",
}

var Values = map[string]int {
    "2": 2,
    "3": 3,
    "4": 4, 
    "5": 5, 
    "6": 6, 
    "7": 7, 
    "8": 8, 
    "9": 9,
    "T": 10,
    "J": 11, 
    "Q": 12,
    "K": 13,
    "A": 14,
}

type Deck []Card

func CreateDeck() Deck {
    var deck Deck
    for _, suit := range Suits {
        for face, value := range Values {
            c := Card{
                Value: value,
                Name: face,
                Suit: suit,
            }
            deck = append(deck, c )
        }
    } 
    return deck
}

func (d *Deck) ShuffleDeck() {
    rand.Shuffle(len(*d), func(i, j int){
        (*d)[i], (*d)[j] = (*d)[j], (*d)[i]
    })
}

func (d *Deck) Deal(n int)[]Card{
    if len(*d) < n {
        return nil
    }
    dealtCards := (*d)[:n]
    (*d) = (*d)[n:]
    return dealtCards

}

