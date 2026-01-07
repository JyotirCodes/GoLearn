package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var value int
    switch card {
        case "ace":
        	value = 11
        case "two":
        	value = 2
        case "three":
        	 value = 3
        case "four":
        	value = 4
        case "five":
        	value = 5
        case "six":
        	value = 6
        case "seven":
        	value = 7
        case "eight":
        	value = 8
        case "nine":
        	value = 9
        case "ten":
        	value = 10
        case "jack":
        	value = 10
        case "queen":
        	value = 10
        case "king":
        	value = 10
        default:
        	value = 0
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var value string
    sumy:= ParseCard(card1) + ParseCard(card2)
    dCard:= ParseCard(dealerCard)
    switch {
        case card1 == "ace" && card2 == "ace":
        	value = "P"
        case (sumy == 21) && dCard < 10:
        	value = "W"
        case sumy == 21 && dCard >= 10:
        	value = "S"
        case sumy >= 17 && sumy <= 20:
        	value = "S"
        case sumy >= 12 && sumy <= 16 && dCard>= 7:
        	value = "H"
        case sumy >= 12 && sumy <= 16 && dCard < 7:
        	value = "S"
        default:
        	value = "H"
    }
    return value
}
