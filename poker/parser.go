package poker

func ExpandHandRange(hr HandRange) []HandType {
	var expanded_range []HandType

	if hr.expansion == "+" {
		if hr.is_pair() {
			for i := hr.get_card1().Rank; i <= rankToOrder["A"]; i++ {
				expanded_range = append(expanded_range, newHandType(newCard(orderToRank[i]), newCard(orderToRank[i]), hr.get_suitedness()))
			}
		} else {
			for i := hr.get_card2().Rank; i <= hr.get_card1().Rank-1; i++ {
				expanded_range = append(expanded_range, newHandType(newCard(orderToRank[hr.get_card1().Rank]), newCard(orderToRank[i]), hr.get_suitedness()))
			}
		}
		return expanded_range
	} else if hr.expansion == "-" { // AKs-ATs
		if hr.is_pair() {
			for i := hr.handType2.card1.Rank; i <= hr.handType1.card1.Rank; i++ {
				expanded_range = append(expanded_range, newHandType(newCard(orderToRank[i]), newCard(orderToRank[i]), hr.get_suitedness()))
			}
		} else {
			for i := hr.handType2.card2.Rank; i <= hr.handType1.card2.Rank; i++ {
				expanded_range = append(expanded_range, newHandType(newCard(orderToRank[hr.get_card1().Rank]), newCard(orderToRank[i]), hr.get_suitedness()))
			}
		}
		return expanded_range
	}

	return []HandType{hr.handType1}
}
