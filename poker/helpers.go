package poker

var rankToOrder = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

var orderToRank = map[int]string{
	1:  "2",
	2:  "3",
	3:  "4",
	4:  "5",
	5:  "6",
	6:  "7",
	7:  "8",
	8:  "9",
	9:  "T",
	10: "J",
	11: "Q",
	12: "K",
	13: "A",
}

func is_pair(hand string) bool {
	return hand[0] == hand[1]
}

func is_suited(hand string) bool {
	return string(hand[2]) == "s"
}
