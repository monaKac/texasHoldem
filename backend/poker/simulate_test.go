package poker

import (
	"testing"
)

func TestSimulate_PocketAcesVsPocketKings(t *testing.T) {
	// AA vs KK pre-flop: AA should win ~81% of the time
	holeCards := []Card{{Hearts, Ace}, {Diamonds, Ace}}
	community := []Card{}

	winPct, _, _ := SimulateWinProbability(holeCards, community, 1, 50000)

	// Allow tolerance: 78% - 85%
	if winPct < 78.0 || winPct > 85.0 {
		t.Errorf("AA vs 1 opponent expected ~81%% win rate, got %.2f%%", winPct)
	}
}

func TestSimulate_WithFlop(t *testing.T) {
	// Player has pair of Aces, flop gives them a set
	holeCards := []Card{{Hearts, Ace}, {Diamonds, Ace}}
	community := []Card{{Clubs, Ace}, {Hearts, Seven}, {Spades, Two}}

	winPct, _, _ := SimulateWinProbability(holeCards, community, 1, 10000)

	// Three aces on the flop should win very often (>90%)
	if winPct < 90.0 {
		t.Errorf("set of Aces on flop should win >90%%, got %.2f%%", winPct)
	}
}

func TestSimulate_SumsTo100(t *testing.T) {
	holeCards := []Card{{Hearts, Ace}, {Diamonds, King}}
	community := []Card{}

	winPct, tiePct, lossPct := SimulateWinProbability(holeCards, community, 1, 10000)

	total := winPct + tiePct + lossPct
	if total < 99.9 || total > 100.1 {
		t.Errorf("win+tie+loss should sum to ~100%%, got %.2f%%", total)
	}
}

func TestSimulate_MultipleOpponents(t *testing.T) {
	// With more opponents, win probability should decrease
	holeCards := []Card{{Hearts, Ace}, {Diamonds, King}}
	community := []Card{}

	win1, _, _ := SimulateWinProbability(holeCards, community, 1, 20000)
	win3, _, _ := SimulateWinProbability(holeCards, community, 3, 20000)

	if win3 >= win1 {
		t.Errorf("more opponents should decrease win rate: 1 opp=%.2f%%, 3 opp=%.2f%%",
			win1, win3)
	}
}

func TestSimulate_RiverAllKnown(t *testing.T) {
	// All 5 community cards known - no randomness in community
	holeCards := []Card{{Hearts, Ace}, {Diamonds, Ace}}
	community := []Card{
		{Clubs, Ace}, {Hearts, Seven}, {Spades, Two},
		{Diamonds, Jack}, {Clubs, Four},
	}

	winPct, _, _ := SimulateWinProbability(holeCards, community, 1, 10000)

	// Three aces with full board should win very often
	if winPct < 95.0 {
		t.Errorf("three aces at river should win >95%%, got %.2f%%", winPct)
	}
}
