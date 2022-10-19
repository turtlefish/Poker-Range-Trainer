package poker

type HandType struct {
	card1      Card
	card2      Card
	suitedness string
}

func (ht HandType) String() string {
	if ht.suitedness == "p" {
		return ht.card1.String() + ht.card2.String()
	} else {
		return ht.card1.String() + ht.card2.String() + ht.suitedness
	}
}

func newHandType(card1 Card, card2 Card, suitedness string) HandType {
	return HandType{card1, card2, suitedness}
}

func newHandTypeFromStr(htStr string) HandType {
	if is_pair(htStr) {
		return HandType{
			newCard(string(htStr[0])),
			newCard(string(htStr[1])),
			"p",
		}
	} else {
		return HandType{
			newCard(string(htStr[0])),
			newCard(string(htStr[1])),
			string(htStr[2]),
		}
	}
}
