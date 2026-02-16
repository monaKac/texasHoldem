package main

import (
	"context"
	"testing"

	pb "github.com/texasholdem/backend/proto"
)

func TestEvaluateHand_RoyalFlush(t *testing.T) {
	s := &server{}
	resp, err := s.EvaluateHand(context.Background(), &pb.EvaluateHandRequest{
		Cards: []string{"HA", "HK", "HQ", "HJ", "HT", "D2", "C3"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.HandRank != "Royal Flush" {
		t.Errorf("expected Royal Flush, got %s", resp.HandRank)
	}
	if len(resp.BestFive) != 5 {
		t.Errorf("expected 5 best cards, got %d", len(resp.BestFive))
	}
}

func TestEvaluateHand_InvalidCardCount(t *testing.T) {
	s := &server{}
	_, err := s.EvaluateHand(context.Background(), &pb.EvaluateHandRequest{
		Cards: []string{"HA", "HK"},
	})
	if err == nil {
		t.Error("expected error for insufficient cards")
	}
}

func TestEvaluateHand_InvalidNotation(t *testing.T) {
	s := &server{}
	_, err := s.EvaluateHand(context.Background(), &pb.EvaluateHandRequest{
		Cards: []string{"HA", "HK", "HQ", "HJ", "HT", "XX", "C3"},
	})
	if err == nil {
		t.Error("expected error for invalid card notation")
	}
}

func TestEvaluateHand_DuplicateCards(t *testing.T) {
	s := &server{}
	_, err := s.EvaluateHand(context.Background(), &pb.EvaluateHandRequest{
		Cards: []string{"HA", "HA", "HQ", "HJ", "HT", "D2", "C3"},
	})
	if err == nil {
		t.Error("expected error for duplicate cards")
	}
}

func TestCompareHands_Player1Wins(t *testing.T) {
	s := &server{}
	resp, err := s.CompareHands(context.Background(), &pb.CompareHandsRequest{
		Player1Cards: []string{"HA", "HK", "HQ", "HJ", "HT", "D2", "C3"},
		Player2Cards: []string{"SA", "SK", "SQ", "SJ", "S9", "D4", "C5"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Winner != 1 {
		t.Errorf("expected Player 1 to win, got winner=%d", resp.Winner)
	}
	if resp.Player1Hand.HandRank != "Royal Flush" {
		t.Errorf("expected Royal Flush for P1, got %s", resp.Player1Hand.HandRank)
	}
}

func TestCompareHands_InvalidInput(t *testing.T) {
	s := &server{}
	_, err := s.CompareHands(context.Background(), &pb.CompareHandsRequest{
		Player1Cards: []string{"HA", "HK"},
		Player2Cards: []string{"SA", "SK", "SQ", "SJ", "S9", "D4", "C5"},
	})
	if err == nil {
		t.Error("expected error for insufficient cards")
	}
}

func TestCalculateWinProbability_Basic(t *testing.T) {
	s := &server{}
	resp, err := s.CalculateWinProbability(context.Background(), &pb.WinProbabilityRequest{
		HoleCards:    []string{"HA", "DA"},
		Community:    []string{},
		NumOpponents: 1,
		Iterations:   10000,
	})
	if err != nil {
		t.Fatal(err)
	}
	total := resp.WinProbability + resp.TieProbability + resp.LossProbability
	if total < 99.0 || total > 101.0 {
		t.Errorf("probabilities should sum to ~100%%, got %.2f%%", total)
	}
}

func TestCalculateWinProbability_InvalidHoleCards(t *testing.T) {
	s := &server{}
	_, err := s.CalculateWinProbability(context.Background(), &pb.WinProbabilityRequest{
		HoleCards:    []string{"HA"},
		Community:    []string{},
		NumOpponents: 1,
		Iterations:   1000,
	})
	if err == nil {
		t.Error("expected error for insufficient hole cards")
	}
}
