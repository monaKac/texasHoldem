package poker

import "testing"

func TestParseCard(t *testing.T) {
	tests := []struct {
		input    string
		wantSuit Suit
		wantRank Rank
		wantErr  bool
	}{
		{"HA", Hearts, Ace, false},
		{"ST", Spades, Ten, false},
		{"D2", Diamonds, Two, false},
		{"CK", Clubs, King, false},
		{"HQ", Hearts, Queen, false},
		{"SJ", Spades, Jack, false},
		{"D9", Diamonds, Nine, false},
		// Error cases
		{"XX", 0, 0, true},
		{"H", 0, 0, true},
		{"HAA", 0, 0, true},
		{"Z5", 0, 0, true},
		{"H1", 0, 0, true},
		{"", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			card, err := ParseCard(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error for %q, got nil", tt.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for %q: %v", tt.input, err)
			}
			if card.Suit != tt.wantSuit {
				t.Errorf("suit: got %c, want %c", card.Suit, tt.wantSuit)
			}
			if card.Rank != tt.wantRank {
				t.Errorf("rank: got %d, want %d", card.Rank, tt.wantRank)
			}
		})
	}
}

func TestCardString(t *testing.T) {
	tests := []struct {
		card Card
		want string
	}{
		{Card{Hearts, Ace}, "HA"},
		{Card{Spades, Ten}, "ST"},
		{Card{Diamonds, Two}, "D2"},
		{Card{Clubs, King}, "CK"},
	}
	for _, tt := range tests {
		got := tt.card.String()
		if got != tt.want {
			t.Errorf("Card{%c, %d}.String() = %q, want %q", tt.card.Suit, tt.card.Rank, got, tt.want)
		}
	}
}

func TestDeck(t *testing.T) {
	deck := Deck()
	if len(deck) != 52 {
		t.Fatalf("deck has %d cards, want 52", len(deck))
	}
	// Check uniqueness
	seen := make(map[string]bool)
	for _, c := range deck {
		s := c.String()
		if seen[s] {
			t.Errorf("duplicate card: %s", s)
		}
		seen[s] = true
	}
}

func TestRemoveCards(t *testing.T) {
	deck := Deck()
	toRemove := []Card{{Hearts, Ace}, {Spades, King}}
	result := RemoveCards(deck, toRemove)
	if len(result) != 50 {
		t.Fatalf("after removing 2 cards, got %d cards, want 50", len(result))
	}
	for _, c := range result {
		if c.Equal(Card{Hearts, Ace}) || c.Equal(Card{Spades, King}) {
			t.Errorf("card %s should have been removed", c.String())
		}
	}
}
