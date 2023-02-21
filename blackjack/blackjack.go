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
	hands := ParseCard(card1) + ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
	case hands == 22:
		return "P"
	case hands == 21 && dealer < 10:
		return "W"
	case hands == 21 && dealer >= 10:
		return "S"
	case hands >= 17 || hands >= 12 && dealer < 7:
		return "S"
	default:
		return "H"
	}
}

/*
case ParseCard("ace") == 11:
	case ParseCard("two") == 2:
	case ParseCard("three") == 3:
	case ParseCard("four") == 4:
	case ParseCard("five") == 5:
	case ParseCard("six") == 6:
	case ParseCard("seven") == 7:
	case ParseCard("eight") == 8:
	case ParseCard("nine") == 9:
	case ParseCard("ten") == 10:
	case ParseCard("jack") == 10:
	case ParseCard("queen") == 10:
	case ParseCard("king") == 10:
case (card1 == "ace") && (card2 == "ace"): // разделение 1 действие
		return "P"
	case (card1 == "ace") && (card2 == "jack" || card2 == "queen" || card2 == "king" || card2 == "ten"):
		return "W"
	case (card1 == "ace") && (card2 == "jack" || card2 == "queen" || card2 == "king" || card2 == "ten") && (dealerCard == "ace" || dealerCard == "five" || dealerCard == "nine" || dealerCard == "queen"): // блекджека нет
		return "S"
	case (card1 == "ten") && (card2 == "seven"): // семнадцать и стоим
		return "S"
	case (card1 == "ten") && (card2 == "eight"): // восемнадцать и стоим
		return "S"
	case (card1 == "ten") && (card2 == "nine"): // девятнадцать и стоим
		return "S"
	case (card1 == "ten") && (card2 == "ten"): // двадцать и стоим
		return "S"
	case (card1 == "ten") && (card2 == "jack"): // двадцать и стоим
		return "S"
	case (card1 == "ten") && (card2 == "queen"): // двадцать и стоим
		return "S"
	case (card1 == "ten") && (card2 == "king"): // двадцать и стоим
		return "S"
	case (card1 == "jack" || card1 == "queen" || card1 == "king" || card1 == "ten") && (card2 == "two" || card2 == "three" || card2 == "four" || card2 == "five" || card2 == "six") && (dealerCard == "seven" || dealerCard == "eight" || dealerCard == "nine" || dealerCard == "ten" || dealerCard == "queen" || dealerCard == "king" || dealerCard == "jack"):
		return "S"
	case (card1 == "jack" || card1 == "queen" || card1 == "king" || card1 == "ten") && (card2 == "two" || card2 == "three" || card2 == "four" || card2 == "five" || card2 == "six"):
		return "H"
	case (card1 == "two") && (card2 == "two" || card2 == "three" || card2 == "four" || card2 == "five" || card2 == "six" || card2 == "seven" || card2 == "eight" || card2 == "nine"):
		return "H"
 && (dealerCard != "ace" || dealerCard != "five" || dealerCard != "nine" || dealerCard != "queen")
case (card1 == "ten" || card1 == "jack" || card1 == "queen" || card1 == "king") && (card2 == "ace"):
		return "W"
	case (card1 == "ten" || card1 == "jack" || card1 == "queen" || card1 == "king") && (card2 == "ace") && (dealerCard == "ten" || dealerCard == "jack" || dealerCard == "queen" || dealerCard == "king" || dealerCard == "ace"):
		return "S"
	case (card1 == "ten") && (card2 == "seven" || card2 == "eight" || card2 == "nine" || card2 == "ten" || card2 == "jack" || card2 == "queen" || card2 == "king"): // стоим 17-20
		return "S"
 || (dealerCard != "ten" && dealerCard != "jack" && dealerCard != "queen" && dealerCard != "king" && dealerCard != "ace")
case (card1 == "ten" || card1 == "jack" || card1 == "queen" || card1 == "king") && (card2 == "ace"): // блекджек
		return "W"
	case (card1 == "ten" || card1 == "jack" || card1 == "queen" || card1 == "king") && (card2 == "ace") && (dealerCard == "ace" || dealerCard == "ten"): // блекджек прогорел
		return "S"
 || (dealerCard != "ace" || dealerCard != "ten")
card == ("ten","jack", "queen", "king")|| card1 == "jack" || card1 == "queen" || card1 == "king"
 || card1 == "jack" || card1 == "queen" || card1 == "king"
 || card2 == "jack" || card2 == "queen" || card2 == "king"
	for value := 0; value < 11; value++ {
		fmt.Println(value)
	}
eight	8
	2	nine	9
three	3	ten	10
four	4	jack	10
five	5	queen	10
six	6	king	10
seven	7	other	0
	case (card1 == "ten" || card2 == "ace") || (dealerCard != "ace"): //блекджек
		return "W"
	case (card1 == "ten" || card2 == "ace") || (dealerCard != "ten"): //блекджек
		return "W"
	case (card1 == "jack" || card2 == "ace") || (dealerCard != "ace"): //блекджек
		return "W"
	case (card1 == "jack" || card2 == "ace") || (dealerCard != "ten"): //блекджек
		return "W"
	case (card1 == "queen" || card2 == "ace") || (dealerCard != "ace"): //блекджек
		return "W"
	case (card1 == "queen" || card2 == "ace") || (dealerCard != "ten"): //блекджек
		return "W"
	case (card1 == "king" || card2 == "ace") || (dealerCard != "ace"): //блекджек
		return "W"
	case (card1 == "king" || card2 == "ace") || (dealerCard != "ten"): //блекджек
		return "S"
	case (card1 == "ten" || card2 == "ace") || (dealerCard == "ace"): //блекджек нет
		return "S"
	case (card1 == "ten" || card2 == "ace") || (dealerCard == "ten"): //блекджек нет
		return "S"
	case (card1 == "jack" || card2 == "ace") || (dealerCard == "ace"): //блекджек нет
		return "S"
	case (card1 == "jack" || card2 == "ace") || (dealerCard == "ten"): //блекджек нет
		return "S"
	case (card1 == "queen" || card2 == "ace") || (dealerCard == "ace"): //блекджек нет
		return "S"
	case (card1 == "queen" || card2 == "ace") || (dealerCard == "ten"): //блекджек нет
		return "S"
	case (card1 == "king" || card2 == "ace") || (dealerCard == "ace"): //блекджекнет
		return "S"
	case (card1 == "king" || card2 == "ace") || (dealerCard == "ten"): //блекджек нет
		return "S"
	case card1 == "ten" || card2 == "seven": // семнадцать и стоим
		return "S"
	case card1 == "ten" || card2 == "eight": // восемнадцать и стоим
		return "S"
	case card1 == "ten" || card2 == "nine": // девятнадцать и стоим
		return "S"
	case card1 == "ten" || card2 == "ten": // двадцать и стоим
		return "S"
	case card1 == "ten" || card2 == "jack": // двадцать и стоим
		return "S"
	case card1 == "ten" || card2 == "queen": // двадцать и стоим
		return "S"
	case card1 == "ten" || card2 == "king": // двадцать и стоим
		return "S"
*/
