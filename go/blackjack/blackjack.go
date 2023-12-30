package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	parsedValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	switch {
  case parsedValue == 22:
		return "P"
  case parsedValue == 21 && dealerValue < 10:
    return "W"
  case parsedValue == 21 && dealerValue >= 10:
    return "S"
  case parsedValue >= 17 && parsedValue <= 20:
    return "S"
  case parsedValue >= 12 && parsedValue <= 16 && dealerValue >= 7:
    return "H"
  case parsedValue >= 12 && parsedValue <= 16 && dealerValue < 7:
    return "S"
  default:
    return "H"
	}
}
