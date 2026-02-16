package poker

import "testing"

func TestCompare_Player1Wins(t *testing.T) {
	p1 := mustParse7(t, "HA", "HK", "HQ", "HJ", "HT", "D2", "C3")
	p2 := mustParse7(t, "H5", "D6", "C7", "S8", "H9", "D2", "C3")

	result, h1, h2 := Compare(p1, p2)
	if result != Player1Wins {
		t.Errorf("expected Player1Wins, got %d (p1=%s, p2=%s)", result, h1.Name, h2.Name)
	}
}

func TestCompare_Player2Wins(t *testing.T) {
	p1 := mustParse7(t, "H8", "D8", "HK", "SQ", "CJ", "D3", "S5")
	p2 := mustParse7(t, "HK", "DK", "CK", "HQ", "DQ", "S2", "C4")

	result, _, _ := Compare(p1, p2)
	if result != Player2Wins {
		t.Errorf("expected Player2Wins, got %d", result)
	}
}

func TestCompare_Tie(t *testing.T) {
	p1 := mustParse7(t, "H2", "D3", "H5", "D6", "C7", "S8", "H9")
	p2 := mustParse7(t, "C2", "S3", "H5", "D6", "C7", "S8", "H9")

	result, _, _ := Compare(p1, p2)
	if result != Tie {
		t.Errorf("expected Tie, got %d", result)
	}
}

func TestCompare_KickerBreaksTie(t *testing.T) {
	p1 := mustParse7(t, "HK", "DK", "HA", "S5", "C3", "D7", "S9")
	p2 := mustParse7(t, "HK", "DK", "HQ", "S5", "C3", "D7", "S9")

	result, _, _ := Compare(p1, p2)
	if result != Player1Wins {
		t.Errorf("expected Player1Wins (ace kicker beats queen kicker), got %d", result)
	}
}

func TestCompare_FullHouseVsFlush(t *testing.T) {
	p1 := mustParse7(t, "HK", "DK", "CK", "HQ", "DQ", "S2", "C4")
	p2 := mustParse7(t, "H2", "H5", "H7", "HJ", "HA", "D3", "C6")

	result, h1, h2 := Compare(p1, p2)
	if result != Player1Wins {
		t.Errorf("Full House should beat Flush: p1=%s, p2=%s", h1.Name, h2.Name)
	}
}

// Comprehensive comparison tests from CSV test matrix.
// Each test: community cards shared, player hole cards differ, expected result.
func TestCompare_CSV(t *testing.T) {
	tests := []struct {
		name      string
		community [5]string
		p1Hole    [2]string
		p2Hole    [2]string
		wantRankP1 HandRank
		wantRankP2 HandRank
		want      CompareResult
	}{
		// --- High Card ---
		{
			name: "HighCard: K kicker > Q kicker",
			community: [5]string{"D6", "S9", "H4", "S3", "C2"},
			p1Hole: [2]string{"SK", "CA"}, p2Hole: [2]string{"HA", "SQ"},
			wantRankP1: HighCard, wantRankP2: HighCard, want: Player1Wins,
		},
		{
			name: "HighCard: same ace-king tie",
			community: [5]string{"D6", "S9", "H4", "S3", "C2"},
			p1Hole: [2]string{"SK", "CA"}, p2Hole: [2]string{"HA", "CK"},
			wantRankP1: HighCard, wantRankP2: HighCard, want: Tie,
		},
		{
			name: "HighCard: Q > J",
			community: [5]string{"D6", "S9", "H4", "H3", "H2"},
			p1Hole: [2]string{"C7", "DQ"}, p2Hole: [2]string{"C8", "DJ"},
			wantRankP1: HighCard, wantRankP2: HighCard, want: Player1Wins,
		},
		// --- One Pair ---
		{
			name: "OnePair: pair K > pair 8",
			community: [5]string{"SK", "HT", "C8", "C7", "D2"},
			p1Hole: [2]string{"DK", "C5"}, p2Hole: [2]string{"H8", "D5"},
			wantRankP1: OnePair, wantRankP2: OnePair, want: Player1Wins,
		},
		{
			name: "OnePair: pair K = pair K same kickers",
			community: [5]string{"SK", "HT", "C8", "C7", "D2"},
			p1Hole: [2]string{"DK", "C4"}, p2Hole: [2]string{"HK", "D5"},
			wantRankP1: OnePair, wantRankP2: OnePair, want: Tie,
		},
		{
			name: "OnePair: pair A kicker 7 > kicker 6",
			community: [5]string{"HA", "DA", "ST", "C9", "D4"},
			p1Hole: [2]string{"D5", "C6"}, p2Hole: [2]string{"H7", "C2"},
			wantRankP1: OnePair, wantRankP2: OnePair, want: Player2Wins,
		},
		// --- Two Pair ---
		{
			name: "TwoPair: A+6 > Q+6",
			community: [5]string{"SA", "DQ", "CK", "D6", "H6"},
			p1Hole: [2]string{"HA", "C3"}, p2Hole: [2]string{"CQ", "H4"},
			wantRankP1: TwoPair, wantRankP2: TwoPair, want: Player1Wins,
		},
		{
			name: "TwoPair: Q+6 = Q+6",
			community: [5]string{"SA", "DQ", "CK", "D6", "H6"},
			p1Hole: [2]string{"HQ", "C3"}, p2Hole: [2]string{"SQ", "H4"},
			wantRankP1: TwoPair, wantRankP2: TwoPair, want: Tie,
		},
		{
			name: "TwoPair: A+K > Q+6",
			community: [5]string{"SA", "DQ", "CK", "D6", "H5"},
			p1Hole: [2]string{"HQ", "C6"}, p2Hole: [2]string{"CA", "HK"},
			wantRankP1: TwoPair, wantRankP2: TwoPair, want: Player2Wins,
		},
		// --- Three of a Kind ---
		{
			name: "ThreeOfAKind: trip J > trip 3",
			community: [5]string{"SA", "D3", "H2", "C8", "SJ"},
			p1Hole: [2]string{"HJ", "DJ"}, p2Hole: [2]string{"C3", "H3"},
			wantRankP1: ThreeOfAKind, wantRankP2: ThreeOfAKind, want: Player1Wins,
		},
		{
			name: "ThreeOfAKind: trip 3 = trip 3",
			community: [5]string{"SA", "D3", "H3", "C8", "SJ"},
			p1Hole: [2]string{"C3", "S2"}, p2Hole: [2]string{"S3", "H2"},
			wantRankP1: ThreeOfAKind, wantRankP2: ThreeOfAKind, want: Tie,
		},
		{
			name: "ThreeOfAKind: trip A kicker K > kicker T",
			community: [5]string{"HA", "SA", "DA", "H3", "HT"},
			p1Hole: [2]string{"S2", "S5"}, p2Hole: [2]string{"H2", "SK"},
			wantRankP1: ThreeOfAKind, wantRankP2: ThreeOfAKind, want: Player2Wins,
		},
		// --- Straight ---
		{
			name: "Straight: 7-high > 6-high",
			community: [5]string{"H3", "S4", "C5", "S6", "HT"},
			p1Hole: [2]string{"D7", "HA"}, p2Hole: [2]string{"H2", "SA"},
			wantRankP1: Straight, wantRankP2: Straight, want: Player1Wins,
		},
		{
			name: "Straight: 7-high = 7-high",
			community: [5]string{"H3", "S4", "C5", "S6", "HT"},
			p1Hole: [2]string{"D7", "HA"}, p2Hole: [2]string{"H7", "SA"},
			wantRankP1: Straight, wantRankP2: Straight, want: Tie,
		},
		{
			name: "Straight: 6-high > wheel (5-high)",
			community: [5]string{"H2", "H3", "S4", "C5", "HT"},
			p1Hole: [2]string{"HA", "S3"}, p2Hole: [2]string{"H6", "SA"},
			wantRankP1: Straight, wantRankP2: Straight, want: Player2Wins,
		},
		// --- Flush ---
		{
			name: "Flush: A-high > Q-high",
			community: [5]string{"D3", "D6", "DT", "C5", "HQ"},
			p1Hole: [2]string{"DK", "DA"}, p2Hole: [2]string{"D2", "DQ"},
			wantRankP1: Flush, wantRankP2: Flush, want: Player1Wins,
		},
		{
			name: "Flush: community flush = community flush",
			community: [5]string{"D3", "D6", "DT", "DJ", "DK"},
			p1Hole: [2]string{"C3", "HA"}, p2Hole: [2]string{"S9", "HJ"},
			wantRankP1: Flush, wantRankP2: Flush, want: Tie,
		},
		{
			name: "Flush: A-high > T-high",
			community: [5]string{"D3", "D6", "DT", "C5", "HQ"},
			p1Hole: [2]string{"D2", "D5"}, p2Hole: [2]string{"DJ", "DA"},
			wantRankP1: Flush, wantRankP2: Flush, want: Player2Wins,
		},
		// --- Full House ---
		{
			name: "FullHouse: QQQ+TT > TTT+QQ",
			community: [5]string{"HQ", "SQ", "HT", "DT", "C3"},
			p1Hole: [2]string{"DQ", "C2"}, p2Hole: [2]string{"CT", "C4"},
			wantRankP1: FullHouse, wantRankP2: FullHouse, want: Player1Wins,
		},
		{
			name: "FullHouse: AAA+QQ = AAA+QQ",
			community: [5]string{"SA", "HQ", "SQ", "HT", "D8"},
			p1Hole: [2]string{"HA", "DQ"}, p2Hole: [2]string{"DA", "CQ"},
			wantRankP1: FullHouse, wantRankP2: FullHouse, want: Tie,
		},
		{
			name: "FullHouse: TTT+QQ < QQQ+TT",
			community: [5]string{"HQ", "SQ", "HT", "DT", "C3"},
			p1Hole: [2]string{"ST", "C2"}, p2Hole: [2]string{"CQ", "C4"},
			wantRankP1: FullHouse, wantRankP2: FullHouse, want: Player2Wins,
		},
		// --- Four of a Kind ---
		{
			name: "FourOfAKind: TTTT+A > TTTT+K",
			community: [5]string{"HT", "ST", "CT", "DT", "HK"},
			p1Hole: [2]string{"HA", "S7"}, p2Hole: [2]string{"DJ", "C5"},
			wantRankP1: FourOfAKind, wantRankP2: FourOfAKind, want: Player1Wins,
		},
		{
			name: "FourOfAKind: 5555+A = 5555+A",
			community: [5]string{"S5", "D5", "C5", "H5", "HA"},
			p1Hole: [2]string{"CT", "HT"}, p2Hole: [2]string{"C4", "SQ"},
			wantRankP1: FourOfAKind, wantRankP2: FourOfAKind, want: Tie,
		},
		{
			name: "FourOfAKind: TTTT+8 < TTTT+K",
			community: [5]string{"HT", "ST", "CT", "DT", "S8"},
			p1Hole: [2]string{"C2", "C3"}, p2Hole: [2]string{"C5", "HK"},
			wantRankP1: FourOfAKind, wantRankP2: FourOfAKind, want: Player2Wins,
		},
		// --- Straight Flush ---
		{
			name: "StraightFlush: 7-high > 6-high",
			community: [5]string{"H3", "H4", "H5", "H6", "HT"},
			p1Hole: [2]string{"H7", "HA"}, p2Hole: [2]string{"H2", "SA"},
			wantRankP1: StraightFlush, wantRankP2: StraightFlush, want: Player1Wins,
		},
		{
			name: "StraightFlush: community SF = community SF",
			community: [5]string{"H3", "H4", "H5", "H6", "H7"},
			p1Hole: [2]string{"HA", "ST"}, p2Hole: [2]string{"CQ", "D6"},
			wantRankP1: StraightFlush, wantRankP2: StraightFlush, want: Tie,
		},
		{
			name: "StraightFlush: T-high < J-high",
			community: [5]string{"S7", "S8", "S9", "ST", "DK"},
			p1Hole: [2]string{"S6", "C2"}, p2Hole: [2]string{"SJ", "D5"},
			wantRankP1: StraightFlush, wantRankP2: StraightFlush, want: Player2Wins,
		},
		// --- Royal Flush ---
		{
			name: "RoyalFlush: community RF = community RF",
			community: [5]string{"DT", "DJ", "DQ", "DK", "DA"},
			p1Hole: [2]string{"H2", "C3"}, p2Hole: [2]string{"S4", "C5"},
			wantRankP1: RoyalFlush, wantRankP2: RoyalFlush, want: Tie,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			community := make([]Card, 5)
			for i, s := range tt.community {
				c, err := ParseCard(s)
				if err != nil {
					t.Fatalf("bad community card %q: %v", s, err)
				}
				community[i] = c
			}

			p1Hole := make([]Card, 2)
			p2Hole := make([]Card, 2)
			for i, s := range tt.p1Hole {
				c, err := ParseCard(s)
				if err != nil {
					t.Fatalf("bad p1 hole card %q: %v", s, err)
				}
				p1Hole[i] = c
			}
			for i, s := range tt.p2Hole {
				c, err := ParseCard(s)
				if err != nil {
					t.Fatalf("bad p2 hole card %q: %v", s, err)
				}
				p2Hole[i] = c
			}

			p1Cards := append(p1Hole, community...)
			p2Cards := append(p2Hole, community...)

			result, h1, h2 := Compare(p1Cards, p2Cards)

			if h1.Rank != tt.wantRankP1 {
				t.Errorf("p1 hand: got %q, want %q", h1.Name, handRankNames[tt.wantRankP1])
			}
			if h2.Rank != tt.wantRankP2 {
				t.Errorf("p2 hand: got %q, want %q", h2.Name, handRankNames[tt.wantRankP2])
			}
			if result != tt.want {
				wantStr := map[CompareResult]string{Player1Wins: "Player1Wins", Player2Wins: "Player2Wins", Tie: "Tie"}
				t.Errorf("got %s, want %s (p1=%s val=%d, p2=%s val=%d)",
					wantStr[result], wantStr[tt.want],
					h1.Name, h1.RankValue, h2.Name, h2.RankValue)
			}
		})
	}
}
