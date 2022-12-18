package blackjack

var m = map[string]int{"ace": 11, "two": 2,
	"three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8,
	"nine": 9, "ten": 10, "jack": 10, "queen": 10, "king": 10}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return m[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
//
//	blackjack_test.go:297: FirstTurn(ten, two, queen) = S, want H
func FirstTurn(card1, card2, dealerCard string) string {
	result := "S"

	yourCardsSum := m[card1] + m[card2]

	if pairOfAces(card1, card2) {
		return "P"
	}

	if blackjack(card1, card2, dealerCard) {
		return "W"
	}
	if yourCardsSum >= 12 && yourCardsSum <= 16 {
		if m[dealerCard] >= 7 {
			return "H"
		} else {
			return "S"
		}
	}

	if yourCardsSum >= 17 && yourCardsSum <= 20 {
		return "S"
	}

	if yourCardsSum <= 11 {
		return "H"
	}

	return result
}

func pairOfAces(card1 string, card2 string) bool {
	return card1 == "ace" && card2 == "ace"
}

func blackjack(card1 string, card2 string, dealerCard string) bool {
	return (m[card1]+m[card2] == 21) && (dealerCard != "ace" && dealerCard != "jack" && dealerCard != "queen" && dealerCard != "king" && dealerCard != "ten")
}
