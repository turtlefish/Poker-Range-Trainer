package poker

import "strings"

type HandRange struct {
	handType1 HandType
	handType2 HandType
	expansion string
}

// func (hr HandRange) String() string {
// 	if &hr.handType2 == nil {
// 		return hr.handType1.String()
// 	} else {
// 		return hr.handType1.String() + hr.expansion + hr.handType2.String()
// 	}
// }

// func newHandRange(handType1 HandType, handType2 HandType, expansion string) HandRange {
// 	return HandRange{handType1, handType2, expansion}
// }

func NewHandRangeFromStr(hrStr string) HandRange {
	var expansion string

	if string(hrStr[len(hrStr)-1]) == "+" {
		expansion = "+"
	} else if strings.Contains(hrStr, "-") {
		expansion = "-"
	} else {
		expansion = ""
	}

	if expansion == "-" {
		hrSplit := strings.Split(hrStr, "-")

		return HandRange{
			newHandTypeFromStr(hrSplit[0]),
			newHandTypeFromStr(hrSplit[1]),
			expansion,
		}
	} else {
		return HandRange{
			newHandTypeFromStr(hrStr),
			HandType{},
			expansion,
		}
	}
}

func (hr HandRange) get_suitedness() string {
	return hr.handType1.suitedness
}

func (hr HandRange) is_pair() bool {
	return hr.handType1.suitedness == "p"
}

func (hr HandRange) get_card1() Card {
	return hr.handType1.card1
}

func (hr HandRange) get_card2() Card {
	return hr.handType1.card2
}
