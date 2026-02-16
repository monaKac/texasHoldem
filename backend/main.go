package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/texasholdem/backend/proto"
	"github.com/texasholdem/backend/poker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const port = ":50051"

type server struct {
	pb.UnimplementedPokerServiceServer
}

func (s *server) EvaluateHand(ctx context.Context, req *pb.EvaluateHandRequest) (*pb.EvaluateHandResponse, error) {
	if len(req.Cards) != 7 {
		return nil, status.Errorf(codes.InvalidArgument, "exactly 7 cards required, got %d", len(req.Cards))
	}

	cards, err := poker.ParseCards(req.Cards)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid card: %v", err)
	}

	if err := validateNoDuplicates(cards); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	result := poker.EvaluateBest7(cards)

	bestFive := make([]string, 5)
	for i, c := range result.BestFive {
		bestFive[i] = c.String()
	}

	return &pb.EvaluateHandResponse{
		HandRank:  result.Name,
		BestFive:  bestFive,
		RankValue: int32(result.RankValue),
	}, nil
}

func (s *server) CompareHands(ctx context.Context, req *pb.CompareHandsRequest) (*pb.CompareHandsResponse, error) {
	if len(req.Player1Cards) != 7 {
		return nil, status.Errorf(codes.InvalidArgument, "player 1 needs exactly 7 cards, got %d", len(req.Player1Cards))
	}
	if len(req.Player2Cards) != 7 {
		return nil, status.Errorf(codes.InvalidArgument, "player 2 needs exactly 7 cards, got %d", len(req.Player2Cards))
	}

	p1Cards, err := poker.ParseCards(req.Player1Cards)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "player 1 invalid card: %v", err)
	}
	p2Cards, err := poker.ParseCards(req.Player2Cards)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "player 2 invalid card: %v", err)
	}

	// Validate no duplicates within each player's hand
	if err := validateNoDuplicates(p1Cards); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "player 1: %v", err)
	}
	if err := validateNoDuplicates(p2Cards); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "player 2: %v", err)
	}

	winner, h1, h2 := poker.Compare(p1Cards, p2Cards)

	bestFive1 := cardsToStrings(h1.BestFive)
	bestFive2 := cardsToStrings(h2.BestFive)

	var description string
	switch winner {
	case poker.Player1Wins:
		description = fmt.Sprintf("Player 1 wins with %s over Player 2's %s", h1.Name, h2.Name)
	case poker.Player2Wins:
		description = fmt.Sprintf("Player 2 wins with %s over Player 1's %s", h2.Name, h1.Name)
	case poker.Tie:
		description = fmt.Sprintf("Tie! Both players have %s", h1.Name)
	}

	return &pb.CompareHandsResponse{
		Winner: int32(winner),
		Player1Hand: &pb.EvaluateHandResponse{
			HandRank:  h1.Name,
			BestFive:  bestFive1,
			RankValue: int32(h1.RankValue),
		},
		Player2Hand: &pb.EvaluateHandResponse{
			HandRank:  h2.Name,
			BestFive:  bestFive2,
			RankValue: int32(h2.RankValue),
		},
		Description: description,
	}, nil
}

func (s *server) CalculateWinProbability(ctx context.Context, req *pb.WinProbabilityRequest) (*pb.WinProbabilityResponse, error) {
	if len(req.HoleCards) != 2 {
		return nil, status.Errorf(codes.InvalidArgument, "exactly 2 hole cards required, got %d", len(req.HoleCards))
	}
	if len(req.Community) > 5 {
		return nil, status.Errorf(codes.InvalidArgument, "at most 5 community cards, got %d", len(req.Community))
	}

	holeCards, err := poker.ParseCards(req.HoleCards)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid hole card: %v", err)
	}

	community, err := poker.ParseCards(req.Community)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid community card: %v", err)
	}

	allCards := append(holeCards, community...)
	if err := validateNoDuplicates(allCards); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	winPct, tiePct, lossPct := poker.SimulateWinProbability(
		holeCards, community, int(req.NumOpponents), int(req.Iterations),
	)

	return &pb.WinProbabilityResponse{
		WinProbability:  winPct,
		TieProbability:  tiePct,
		LossProbability: lossPct,
		IterationsRun:   req.Iterations,
	}, nil
}

func validateNoDuplicates(cards []poker.Card) error {
	seen := make(map[string]bool)
	for _, c := range cards {
		s := c.String()
		if seen[s] {
			return fmt.Errorf("duplicate card: %s", s)
		}
		seen[s] = true
	}
	return nil
}

func cardsToStrings(cards [5]poker.Card) []string {
	strs := make([]string, 5)
	for i, c := range cards {
		strs[i] = c.String()
	}
	return strs
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPokerServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)
	log.Printf("gRPC server listening on %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
