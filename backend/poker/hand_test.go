package poker

import "testing"

// Helper to create a [5]Card from notation strings.
func mustParse5(t *testing.T, strs ...string) [5]Card {
	t.Helper()
	if len(strs) != 5 {
		t.Fatalf("need exactly 5 cards, got %d", len(strs))
	}
	var cards [5]Card
	for i, s := range strs {
		c, err := ParseCard(s)
		if err != nil {
			t.Fatalf("invalid card %q: %v", s, err)
		}
		cards[i] = c
	}
	return cards
}

func mustParse7(t *testing.T, strs ...string) []Card {
	t.Helper()
	if len(strs) != 7 {
		t.Fatalf("need exactly 7 cards, got %d", len(strs))
	}
	cards := make([]Card, 7)
	for i, s := range strs {
		c, err := ParseCard(s)
		if err != nil {
			t.Fatalf("invalid card %q: %v", s, err)
		}
		cards[i] = c
	}
	return cards
}

func TestEvaluate5_AllHandRanks(t *testing.T) {
	tests := []struct {
		name     string
		cards    [5]string
		wantRank HandRank
	}{
		{
			name:     "Royal Flush",
			cards:    [5]string{"HA", "HK", "HQ", "HJ", "HT"},
			wantRank: RoyalFlush,
		},
		{
			name:     "Straight Flush",
			cards:    [5]string{"S9", "S8", "S7", "S6", "S5"},
			wantRank: StraightFlush,
		},
		{
			name:     "Four of a Kind",
			cards:    [5]string{"HA", "DA", "CA", "SA", "HK"},
			wantRank: FourOfAKind,
		},
		{
			name:     "Full House",
			cards:    [5]string{"HK", "DK", "CK", "HQ", "DQ"},
			wantRank: FullHouse,
		},
		{
			name:     "Flush",
			cards:    [5]string{"H2", "H5", "H7", "HJ", "HA"},
			wantRank: Flush,
		},
		{
			name:     "Straight",
			cards:    [5]string{"H5", "D6", "C7", "S8", "H9"},
			wantRank: Straight,
		},
		{
			name:     "Three of a Kind",
			cards:    [5]string{"HJ", "DJ", "CJ", "H3", "S7"},
			wantRank: ThreeOfAKind,
		},
		{
			name:     "Two Pair",
			cards:    [5]string{"HA", "DA", "HK", "DK", "S3"},
			wantRank: TwoPair,
		},
		{
			name:     "One Pair",
			cards:    [5]string{"H8", "D8", "HK", "SQ", "CJ"},
			wantRank: OnePair,
		},
		{
			name:     "High Card",
			cards:    [5]string{"HA", "DK", "CQ", "SJ", "H9"},
			wantRank: HighCard,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards := mustParse5(t, tt.cards[0], tt.cards[1], tt.cards[2], tt.cards[3], tt.cards[4])
			result := Evaluate5(cards)
			if result.Rank != tt.wantRank {
				t.Errorf("got %q (%d), want %q (%d)", result.Name, result.Rank, handRankNames[tt.wantRank], tt.wantRank)
			}
		})
	}
}

func TestEvaluate5_WheelStraight(t *testing.T) {
	// A-2-3-4-5: the lowest straight, Ace acts as 1
	cards := mustParse5(t, "HA", "D2", "C3", "S4", "H5")
	result := Evaluate5(cards)
	if result.Rank != Straight {
		t.Errorf("wheel should be Straight, got %q", result.Name)
	}

	// A-2-3-4-5 should be lower than 2-3-4-5-6
	higher := mustParse5(t, "H2", "D3", "C4", "S5", "H6")
	higherResult := Evaluate5(higher)
	if higherResult.Rank != Straight {
		t.Fatalf("expected Straight, got %q", higherResult.Name)
	}
	if result.RankValue >= higherResult.RankValue {
		t.Errorf("wheel (A-5) should rank lower than 6-high straight: wheel=%d, six-high=%d",
			result.RankValue, higherResult.RankValue)
	}
}

func TestEvaluate5_WheelStraightFlush(t *testing.T) {
	cards := mustParse5(t, "HA", "H2", "H3", "H4", "H5")
	result := Evaluate5(cards)
	if result.Rank != StraightFlush {
		t.Errorf("A-2-3-4-5 suited should be Straight Flush, got %q", result.Name)
	}
}

func TestEvaluate5_AceHighStraight(t *testing.T) {
	cards := mustParse5(t, "HT", "DJ", "CQ", "SK", "HA")
	result := Evaluate5(cards)
	if result.Rank != Straight {
		t.Errorf("T-J-Q-K-A should be Straight, got %q", result.Name)
	}
}

func TestEvaluate5_AceHighBeatsWheelStraight(t *testing.T) {
	wheel := mustParse5(t, "HA", "D2", "C3", "S4", "H5")
	aceHigh := mustParse5(t, "HT", "DJ", "CQ", "SK", "DA")

	wheelResult := Evaluate5(wheel)
	aceHighResult := Evaluate5(aceHigh)

	if aceHighResult.RankValue <= wheelResult.RankValue {
		t.Errorf("ace-high straight should beat wheel: ace-high=%d, wheel=%d",
			aceHighResult.RankValue, wheelResult.RankValue)
	}
}

func TestEvaluate5_RankOrdering(t *testing.T) {
	// Verify the fundamental ordering: Royal Flush > Straight Flush > ... > High Card
	hands := []struct {
		name  string
		cards [5]string
	}{
		{"High Card", [5]string{"HA", "DK", "CQ", "SJ", "H9"}},
		{"One Pair", [5]string{"H8", "D8", "HK", "SQ", "CJ"}},
		{"Two Pair", [5]string{"HA", "DA", "HK", "DK", "S3"}},
		{"Three of a Kind", [5]string{"HJ", "DJ", "CJ", "H3", "S7"}},
		{"Straight", [5]string{"H5", "D6", "C7", "S8", "H9"}},
		{"Flush", [5]string{"H2", "H5", "H7", "HJ", "HA"}},
		{"Full House", [5]string{"HK", "DK", "CK", "HQ", "DQ"}},
		{"Four of a Kind", [5]string{"HA", "DA", "CA", "SA", "HK"}},
		{"Straight Flush", [5]string{"S9", "S8", "S7", "S6", "S5"}},
		{"Royal Flush", [5]string{"HA", "HK", "HQ", "HJ", "HT"}},
	}

	var prevValue int64
	for i, h := range hands {
		cards := mustParse5(t, h.cards[0], h.cards[1], h.cards[2], h.cards[3], h.cards[4])
		result := Evaluate5(cards)
		if i > 0 && result.RankValue <= prevValue {
			t.Errorf("%s (value=%d) should rank higher than %s (value=%d)",
				h.name, result.RankValue, hands[i-1].name, prevValue)
		}
		prevValue = result.RankValue
	}
}

func TestEvaluateBest7_FindsRoyalFlush(t *testing.T) {
	cards := mustParse7(t, "HA", "HK", "HQ", "HJ", "HT", "D2", "C3")
	result := EvaluateBest7(cards)
	if result.Rank != RoyalFlush {
		t.Errorf("expected Royal Flush, got %q", result.Name)
	}
}

func TestEvaluateBest7_FindsBestHand(t *testing.T) {
	// 7 cards contain both a flush and a full house; full house should win
	cards := mustParse7(t, "HK", "DK", "CK", "HQ", "DQ", "H5", "H3")
	result := EvaluateBest7(cards)
	if result.Rank != FullHouse {
		t.Errorf("expected Full House, got %q", result.Name)
	}
}

func TestEvaluate5_KickerComparison(t *testing.T) {
	// Pair of Kings, Ace kicker vs Pair of Kings, Queen kicker
	pairKingsAce := mustParse5(t, "HK", "DK", "HA", "S5", "C3")
	pairKingsQueen := mustParse5(t, "HK", "DK", "HQ", "S5", "C3")

	r1 := Evaluate5(pairKingsAce)
	r2 := Evaluate5(pairKingsQueen)

	if r1.Rank != OnePair || r2.Rank != OnePair {
		t.Fatalf("both should be One Pair, got %q and %q", r1.Name, r2.Name)
	}
	if r1.RankValue <= r2.RankValue {
		t.Errorf("pair of Kings with Ace kicker should beat Queen kicker: %d vs %d",
			r1.RankValue, r2.RankValue)
	}
}

func TestEvaluate5_TwoPairKicker(t *testing.T) {
	// Two pair (Aces and Kings) with Queen kicker vs Jack kicker
	tpQueen := mustParse5(t, "HA", "DA", "HK", "DK", "CQ")
	tpJack := mustParse5(t, "HA", "DA", "HK", "DK", "CJ")

	r1 := Evaluate5(tpQueen)
	r2 := Evaluate5(tpJack)

	if r1.RankValue <= r2.RankValue {
		t.Errorf("two pair with Queen kicker should beat Jack kicker")
	}
}

func TestEvaluate5_FullHouseTiebreak(t *testing.T) {
	// Full house: Kings full of Queens vs Kings full of Jacks
	kingsQueens := mustParse5(t, "HK", "DK", "CK", "HQ", "DQ")
	kingsJacks := mustParse5(t, "HK", "DK", "CK", "HJ", "DJ")

	r1 := Evaluate5(kingsQueens)
	r2 := Evaluate5(kingsJacks)

	if r1.Rank != FullHouse || r2.Rank != FullHouse {
		t.Fatal("both should be Full House")
	}
	if r1.RankValue <= r2.RankValue {
		t.Errorf("Kings full of Queens should beat Kings full of Jacks")
	}
}
