package poker

import "fmt"

// Rank represents a card rank as an integer 2-14 (14 = Ace).
type Rank int

const (
	Two   Rank = 2
	Three Rank = 3
	Four  Rank = 4
	Five  Rank = 5
	Six   Rank = 6
	Seven Rank = 7
	Eight Rank = 8
	Nine  Rank = 9
	Ten   Rank = 10
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
	Ace   Rank = 14
)

// Suit represents a card suit.
type Suit byte

const (
	Hearts   Suit = 'H'
	Diamonds Suit = 'D'
	Clubs    Suit = 'C'
	Spades   Suit = 'S'
)

// Card holds a suit and rank.
type Card struct {
	Suit Suit
	Rank Rank
}

var rankToChar = map[Rank]byte{
	Two: '2', Three: '3', Four: '4', Five: '5', Six: '6',
	Seven: '7', Eight: '8', Nine: '9', Ten: 'T',
	Jack: 'J', Queen: 'Q', King: 'K', Ace: 'A',
}

var charToRank = map[byte]Rank{
	'2': Two, '3': Three, '4': Four, '5': Five, '6': Six,
	'7': Seven, '8': Eight, '9': Nine, 'T': Ten,
	'J': Jack, 'Q': Queen, 'K': King, 'A': Ace,
}

// ParseCard parses a 2-char notation string (e.g. "HA") into a Card.
// Format: suit (H/D/C/S) + rank (2-9/T/J/Q/K/A)
func ParseCard(s string) (Card, error) {
	if len(s) != 2 {
		return Card{}, fmt.Errorf("card notation must be exactly 2 characters, got %q", s)
	}

	suit := Suit(s[0])
	switch suit {
	case Hearts, Diamonds, Clubs, Spades:
	default:
		return Card{}, fmt.Errorf("invalid suit %q, must be H/D/C/S", s[0])
	}

	rank, ok := charToRank[s[1]]
	if !ok {
		return Card{}, fmt.Errorf("invalid rank %q, must be 2-9/T/J/Q/K/A", s[1])
	}

	return Card{Suit: suit, Rank: rank}, nil
}

// ParseCards parses a slice of card notation strings.
func ParseCards(strs []string) ([]Card, error) {
	cards := make([]Card, len(strs))
	for i, s := range strs {
		c, err := ParseCard(s)
		if err != nil {
			return nil, fmt.Errorf("card %d: %w", i, err)
		}
		cards[i] = c
	}
	return cards, nil
}

// String returns the 2-char notation for a Card.
func (c Card) String() string {
	r, ok := rankToChar[c.Rank]
	if !ok {
		return "??"
	}
	return string([]byte{byte(c.Suit), r})
}

// Equal returns true if two cards have the same suit and rank.
func (c Card) Equal(other Card) bool {
	return c.Suit == other.Suit && c.Rank == other.Rank
}

// Deck returns all 52 cards.
func Deck() []Card {
	suits := []Suit{Hearts, Diamonds, Clubs, Spades}
	ranks := []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	deck := make([]Card, 0, 52)
	for _, s := range suits {
		for _, r := range ranks {
			deck = append(deck, Card{Suit: s, Rank: r})
		}
	}
	return deck
}

// RemoveCards returns a new slice with the specified cards removed.
func RemoveCards(deck []Card, toRemove []Card) []Card {
	result := make([]Card, 0, len(deck))
	for _, c := range deck {
		found := false
		for _, r := range toRemove {
			if c.Equal(r) {
				found = true
				break
			}
		}
		if !found {
			result = append(result, c)
		}
	}
	return result
}
