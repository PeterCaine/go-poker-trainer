package poker

import "sort"

type HandRanking int 

const (
    HighCard HandRanking = iota
    OnePair
    TwoPair
    ThreeOfAKind
    Straight
    Flush
    FullHouse
    FourOfAKind
    StraightFlush
    RoyalFlush
)

func equalSlice(a, b []int) bool {
    if len(a) != len(b){
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func isSequential(values []int) bool {
    for i:=1; i<len(values); i++ {
        if values[i] != values[i-1]+1 {
            break
        } 
        if i == len(values)-1{
            return true
        }
    }
    lowStraight := []int{1,2,3,4,14}
    if len(values) == 5 && equalSlice(values, lowStraight){
        return true
    }
    return false
}

func hasCount(valueCount map[int]int, count int) bool {
    for _, v := range valueCount{
        if v == count {
            return true
        }
    }
    return false
}

func hasFullHouse(valueCount map[int]int) bool {
    if hasCount(valueCount, 2) && hasCount(valueCount, 3) {
        return true
    }
    return false
}

func hasTwoPair(valueCount map[int]int) bool{
    pairCount := 0
    for _, val := range valueCount{
        if val ==2 {
            pairCount++
        }
    }
    return pairCount == 2
}

func EvaluateHand(hand []Card) HandRanking {
    values := []int{}
    suits := map[string]int{}
    valueCount := map[int]int{}

    for _, card := range hand {
        values = append(values, card.Value)
        suits[card.Suit]++
        valueCount[card.Value]++
    }

    sort.Ints(values)

    isFlush := len(suits) == 1
    isStraight := isSequential(values)

    switch {
    case isFlush && isStraight && values[len(values)-1] == 14:
        return RoyalFlush
    case isFlush && isStraight:
        return StraightFlush
    case hasCount(valueCount, 4):
        return FourOfAKind
    case hasFullHouse(valueCount):
        return FullHouse
    case isFlush:
        return Flush
    case isStraight:
        return Straight
    case hasCount(valueCount, 3):
        return ThreeOfAKind
    case hasTwoPair(valueCount):
        return TwoPair
    case hasCount(valueCount, 2):
        return OnePair
    default:
        return HighCard
    }


}

func handRankToString(rank HandRanking, bestHand []Card) string {
	switch rank {
	case RoyalFlush:
		return "Royal Flush"
	case StraightFlush:
		return "Straight Flush"
	case FourOfAKind:
		return "Four of a Kind"
	case FullHouse:
		return "Full House"
	case Flush:
		return "Flush"
	case Straight:
		return "Straight"
	case ThreeOfAKind:
		return "Three of a Kind"
	case TwoPair:
		return "Two Pair"
	case OnePair:
		return "One Pair"
	default:
		// If High Card, display the highest card in hand
		sort.Slice(bestHand, func(i, j int) bool {
			return bestHand[i].Value > bestHand[j].Value
		})
		return "High Card: " + bestHand[0].Name
	}
}

func FindBestHand(communityCards []Card, playerHand []Card) string {
    // combine community cards and player hand to make 7 communityCards
    cards := append(communityCards, playerHand...)
    combinations := allFiveCardCombinations(cards)

    bestRank := HighCard
    var bestHand []Card

    for _, hand := range combinations {
        rank := EvaluateHand(hand)
        if rank > bestRank {
            bestRank = rank
            bestHand = hand
        }
    }


    return handRankToString(bestRank, bestHand)
}

func allFiveCardCombinations(cards []Card )[][]Card{
    n:=len(cards)
    var combinations [][]Card
    for i:=0; i<n; i++{
        for j:=i+1; j<n; j++{
            for k:=j+1; k<n; k++{
                for l:=k+1; k<n; k++{
                    for m:=l+1; m<n; l++ {
                        hand := []Card{cards[i], cards[j], cards[k], cards[l], cards[m]}
                        combinations = append(combinations, hand)
                    }
                }
            }
        }
    }
    return combinations
}
