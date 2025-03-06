package poker

type Game struct {
    Deck Deck
    PlayerHand []Card
    CommunityCards []Card
    CurrentPhase string
}

func NewGame() *Game {
    g := &Game {
        Deck: CreateDeck(),
        PlayerHand: []Card{},
        CommunityCards: []Card{},
        CurrentPhase: "preflop",
    }
    g.Deck.ShuffleDeck()
    g.PlayerHand = g.Deck.Deal(2)
    return g

}

func (g *Game) DealNextPhase(){
    switch g.CurrentPhase {
    case "preflop":
        g.CommunityCards = g.Deck.Deal(3)
        g.CurrentPhase = "flop"
    case "flop":
        g.CommunityCards = append(g.CommunityCards, g.Deck.Deal(1)[0])
        g.CurrentPhase = "turn"
    case "turn":
        g.CommunityCards = append(g.CommunityCards, g.Deck.Deal(1)[0])
        g.CurrentPhase = "river"
    case "river":
        g.CurrentPhase = "showdown"
    }
}
