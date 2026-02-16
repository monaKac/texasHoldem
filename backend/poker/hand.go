package poker

import "sort"

// HandRank enumerates hand rankings from worst (0) to best (9).
type HandRank int

const (
	HighCard      HandRank = 0
	OnePair       HandRank = 1
	TwoPair       HandRank = 2
	ThreeOfAKind  HandRank = 3
	Straight      HandRank = 4
	Flush         HandRank = 5
	FullHouse     HandRank = 6
	FourOfAKind   HandRank = 7
	StraightFlush HandRank = 8
	RoyalFlush    HandRank = 9
)

var handRankNames = map[HandRank]string{
	HighCard:      "High Card",
	OnePair:       "One Pair",
	TwoPair:       "Two Pair",
	ThreeOfAKind:  "Three of a Kind",
	Straight:      "Straight",
	Flush:         "Flush",
	FullHouse:     "Full House",
	FourOfAKind:   "Four of a Kind",
	StraightFlush: "Straight Flush",
	RoyalFlush:    "Royal Flush",
}

// HandResult holds the evaluation result for a hand.
type HandResult struct {
	Rank      HandRank
	RankValue int64    // Composite numeric score for ordering
	BestFive  [5]Card  // The 5 cards forming the best hand
	Name      string   // Human-readable name
}

// Evaluate5 evaluates exactly 5 cards and returns a HandResult.
// This is the core Norvig-style evaluation.
func Evaluate5(cards [5]Card) HandResult {
	// Sort by rank descending
	sorted := cards
	sort.Slice(sorted[:], func(i, j int) bool {
		return sorted[i].Rank > sorted[j].Rank
	})

	// Count ranks and suits
	rankCounts := make(map[Rank]int)
	suitCounts := make(map[Suit]int)
	for _, c := range sorted {
		rankCounts[c.Rank]++
		suitCounts[c.Suit]++
	}

	isFlush := len(suitCounts) == 1

	// Check for straight
	isStraight := false
	isWheel := false // A-2-3-4-5
	ranks := sortedRanksDesc(sorted)

	if ranks[0]-ranks[4] == 4 && len(rankCounts) == 5 {
		isStraight = true
	}
	// Special case: wheel (A-5-4-3-2)
	if ranks[0] == int(Ace) && ranks[1] == int(Five) && ranks[2] == int(Four) && ranks[3] == int(Three) && ranks[4] == int(Two) {
		isStraight = true
		isWheel = true
	}

	// Classify hand
	groups := groupByCount(rankCounts)
	var hr HandRank
	var tiebreakers []int

	switch {
	case isFlush && isStraight && !isWheel && ranks[0] == int(Ace):
		hr = RoyalFlush
		tiebreakers = ranks

	case isFlush && isStraight:
		hr = StraightFlush
		if isWheel {
			tiebreakers = []int{5, 4, 3, 2, 1} // Ace counts as 1
		} else {
			tiebreakers = ranks
		}

	case groups[0].count == 4:
		hr = FourOfAKind
		tiebreakers = append([]int{int(groups[0].rank)}, int(groups[1].rank))

	case groups[0].count == 3 && groups[1].count == 2:
		hr = FullHouse
		tiebreakers = []int{int(groups[0].rank), int(groups[1].rank)}

	case isFlush:
		hr = Flush
		tiebreakers = ranks

	case isStraight:
		hr = Straight
		if isWheel {
			tiebreakers = []int{5, 4, 3, 2, 1}
		} else {
			tiebreakers = ranks
		}

	case groups[0].count == 3:
		hr = ThreeOfAKind
		tiebreakers = []int{int(groups[0].rank), int(groups[1].rank), int(groups[2].rank)}

	case groups[0].count == 2 && groups[1].count == 2:
		hr = TwoPair
		tiebreakers = []int{int(groups[0].rank), int(groups[1].rank), int(groups[2].rank)}

	case groups[0].count == 2:
		hr = OnePair
		tiebreakers = []int{int(groups[0].rank), int(groups[1].rank), int(groups[2].rank), int(groups[3].rank)}

	default:
		hr = HighCard
		tiebreakers = ranks
	}

	rankValue := computeRankValue(hr, tiebreakers)

	return HandResult{
		Rank:      hr,
		RankValue: rankValue,
		BestFive:  sorted,
		Name:      handRankNames[hr],
	}
}

// EvaluateBest7 takes 7 cards and returns the best 5-card hand
// by checking all C(7,5) = 21 combinations.
func EvaluateBest7(cards []Card) HandResult {
	var best HandResult
	first := true

	// Choose 2 cards to skip
	for i := 0; i < 7; i++ {
		for j := i + 1; j < 7; j++ {
			var five [5]Card
			idx := 0
			for k := 0; k < 7; k++ {
				if k != i && k != j {
					five[idx] = cards[k]
					idx++
				}
			}
			result := Evaluate5(five)
			if first || result.RankValue > best.RankValue {
				best = result
				first = false
			}
		}
	}
	return best
}

// rankGroup is a rank with its count, used for classification.
type rankGroup struct {
	rank  Rank
	count int
}

// groupByCount groups ranks by count (descending count, then descending rank).
func groupByCount(rankCounts map[Rank]int) []rankGroup {
	groups := make([]rankGroup, 0, len(rankCounts))
	for r, c := range rankCounts {
		groups = append(groups, rankGroup{rank: r, count: c})
	}
	sort.Slice(groups, func(i, j int) bool {
		if groups[i].count != groups[j].count {
			return groups[i].count > groups[j].count
		}
		return groups[i].rank > groups[j].rank
	})
	return groups
}

// sortedRanksDesc returns the ranks of the sorted cards as ints, descending.
func sortedRanksDesc(cards [5]Card) []int {
	ranks := make([]int, 5)
	for i, c := range cards {
		ranks[i] = int(c.Rank)
	}
	return ranks
}

// computeRankValue creates a single int64 for hand comparison.
// Format: HandRank * 15^5 + tb[0] * 15^4 + tb[1] * 15^3 + ...
func computeRankValue(hr HandRank, tiebreakers []int) int64 {
	const base = 15
	value := int64(hr) * pow(base, 5)
	for i, tb := range tiebreakers {
		if i >= 5 {
			break
		}
		value += int64(tb) * pow(base, 4-i)
	}
	return value
}

func pow(base, exp int) int64 {
	result := int64(1)
	b := int64(base)
	for i := 0; i < exp; i++ {
		result *= b
	}
	return result
}
