package poker

import (
	"math/rand"
)

const defaultIterations = 10000

// SimulateWinProbability runs Monte Carlo simulation for win probability.
// holeCards: player's 2 hole cards.
// community: 0-5 known community cards.
// numOpponents: number of opponents (each gets 2 random hole cards).
// iterations: number of simulations to run (0 = default 10000).
func SimulateWinProbability(holeCards []Card, community []Card, numOpponents, iterations int) (winPct, tiePct, lossPct float64) {
	if iterations <= 0 {
		iterations = defaultIterations
	}
	if numOpponents <= 0 {
		numOpponents = 1
	}

	knownCards := make([]Card, 0, len(holeCards)+len(community))
	knownCards = append(knownCards, holeCards...)
	knownCards = append(knownCards, community...)

	remainingBase := RemoveCards(Deck(), knownCards)
	communityNeeded := 5 - len(community)

	wins, ties, losses := 0, 0, 0

	for i := 0; i < iterations; i++ {
		remaining := make([]Card, len(remainingBase))
		copy(remaining, remainingBase)

		// Fisher-Yates shuffle
		for j := len(remaining) - 1; j > 0; j-- {
			k := rand.Intn(j + 1)
			remaining[j], remaining[k] = remaining[k], remaining[j]
		}

		drawIdx := 0

		// Complete community cards
		fullCommunity := make([]Card, 5)
		copy(fullCommunity, community)
		for j := len(community); j < 5; j++ {
			fullCommunity[j] = remaining[drawIdx]
			drawIdx++
		}

		// Build player's 7-card hand
		playerHand := make([]Card, 7)
		copy(playerHand[:2], holeCards)
		copy(playerHand[2:], fullCommunity)
		playerResult := EvaluateBest7(playerHand)

		// Evaluate all opponents
		playerWins := true
		isTie := false
		for opp := 0; opp < numOpponents; opp++ {
			oppHole := remaining[drawIdx : drawIdx+2]
			drawIdx += 2

			oppHand := make([]Card, 7)
			copy(oppHand[:2], oppHole)
			copy(oppHand[2:], fullCommunity)
			oppResult := EvaluateBest7(oppHand)

			if oppResult.RankValue > playerResult.RankValue {
				playerWins = false
				isTie = false
				break
			} else if oppResult.RankValue == playerResult.RankValue {
				isTie = true
			}
		}

		if !playerWins {
			losses++
		} else if isTie {
			ties++
		} else {
			wins++
		}

		_ = communityNeeded // used indirectly above
	}

	winPct = float64(wins) / float64(iterations) * 100.0
	tiePct = float64(ties) / float64(iterations) * 100.0
	lossPct = float64(losses) / float64(iterations) * 100.0
	return
}
