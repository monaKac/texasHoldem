package poker

// CompareResult indicates who won.
type CompareResult int

const (
	Tie         CompareResult = 0
	Player1Wins CompareResult = 1
	Player2Wins CompareResult = 2
)

// Compare evaluates both players' best hands from 7 cards and returns the winner.
func Compare(player1Cards, player2Cards []Card) (CompareResult, HandResult, HandResult) {
	h1 := EvaluateBest7(player1Cards)
	h2 := EvaluateBest7(player2Cards)

	if h1.RankValue > h2.RankValue {
		return Player1Wins, h1, h2
	} else if h2.RankValue > h1.RankValue {
		return Player2Wins, h1, h2
	}
	return Tie, h1, h2
}
