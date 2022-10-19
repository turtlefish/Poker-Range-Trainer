package poker

type Card struct {
	Rank int
}

func (c Card) String() string {
	return orderToRank[c.Rank]
}

func newCard(rank string) Card {
	return Card{rankToOrder[rank]}
}
