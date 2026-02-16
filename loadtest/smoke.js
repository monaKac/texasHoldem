import grpc from 'k6/net/grpc';
import { check, sleep } from 'k6';
import { Trend, Counter } from 'k6/metrics';

// Custom metrics
const evaluateDuration = new Trend('evaluate_duration', true);
const compareDuration = new Trend('compare_duration', true);
const probabilityDuration = new Trend('probability_duration', true);
const grpcErrors = new Counter('grpc_errors');

const client = new grpc.Client();
client.load(['../proto'], 'poker.proto');

// Load test stages: ramp up → sustain → spike → sustain → ramp down
export const options = {
  stages: [
    { duration: '10s', target: 50 },   // ramp up to 50 VUs
    { duration: '30s', target: 50 },   // sustain 50 VUs
    { duration: '10s', target: 100 },  // spike to 100 VUs
    { duration: '30s', target: 100 },  // sustain 100 VUs
    { duration: '10s', target: 0 },    // ramp down
  ],
  thresholds: {
    grpc_req_duration: ['p(95)<2000'],       // 95th percentile under 2s
    evaluate_duration: ['p(95)<500'],         // evaluate under 500ms
    compare_duration: ['p(95)<500'],          // compare under 500ms
    probability_duration: ['p(95)<3000'],     // probability under 3s (Monte Carlo is heavier)
    grpc_errors: ['count<10'],               // fewer than 10 gRPC errors
  },
};

// Sample card combinations for realistic testing (no duplicate cards within a set)
const HANDS = [
  { hole: ['HA', 'HK'], community: ['HQ', 'HJ', 'HT', 'D2', 'C3'] },  // Royal flush
  { hole: ['S9', 'S8'], community: ['S7', 'S6', 'S5', 'D2', 'C3'] },   // Straight flush
  { hole: ['HA', 'DA'], community: ['CA', 'SA', 'HK', 'D2', 'C3'] },   // Four of a kind
  { hole: ['HK', 'DK'], community: ['CK', 'HQ', 'DQ', 'S2', 'C4'] },  // Full house
  { hole: ['H2', 'H5'], community: ['H7', 'HJ', 'HA', 'D3', 'C6'] },  // Flush
  { hole: ['H5', 'D6'], community: ['C7', 'S8', 'H9', 'D2', 'C3'] },   // Straight
  { hole: ['HJ', 'DJ'], community: ['CJ', 'H3', 'S7', 'D5', 'C9'] },  // Three of a kind
  { hole: ['HA', 'DA'], community: ['HK', 'DK', 'S3', 'C7', 'D9'] },   // Two pair
  { hole: ['H8', 'D8'], community: ['HK', 'SQ', 'CJ', 'D3', 'S5'] },  // One pair
  { hole: ['HA', 'DK'], community: ['CQ', 'SJ', 'H9', 'D3', 'C5'] },  // High card
];

// Pre-built comparison scenarios from CSV test matrix (no duplicate cards)
const COMPARISONS = [
  // High Card
  { p1: ['SK', 'CA'], p2: ['HA', 'SQ'], community: ['D6', 'S9', 'H4', 'S3', 'C2'] },
  { p1: ['C7', 'DQ'], p2: ['C8', 'DJ'], community: ['D6', 'S9', 'H4', 'H3', 'H2'] },
  // One Pair
  { p1: ['DK', 'C5'], p2: ['H8', 'D5'], community: ['SK', 'HT', 'C8', 'C7', 'D2'] },
  { p1: ['D5', 'C6'], p2: ['H7', 'C2'], community: ['HA', 'DA', 'ST', 'C9', 'D4'] },
  // Two Pair
  { p1: ['HA', 'C3'], p2: ['CQ', 'H4'], community: ['SA', 'DQ', 'CK', 'D6', 'H6'] },
  { p1: ['HQ', 'C6'], p2: ['CA', 'HK'], community: ['SA', 'DQ', 'CK', 'D6', 'H5'] },
  // Three of a Kind
  { p1: ['HJ', 'DJ'], p2: ['C3', 'H3'], community: ['SA', 'D3', 'H2', 'C8', 'SJ'] },
  { p1: ['S2', 'S5'], p2: ['H2', 'SK'], community: ['HA', 'SA', 'DA', 'H3', 'HT'] },
  // Straight
  { p1: ['D7', 'HA'], p2: ['H2', 'SA'], community: ['H3', 'S4', 'C5', 'S6', 'HT'] },
  { p1: ['HA', 'S3'], p2: ['H6', 'SA'], community: ['H2', 'H3', 'S4', 'C5', 'HT'] },
  // Flush
  { p1: ['DK', 'DA'], p2: ['D2', 'DQ'], community: ['D3', 'D6', 'DT', 'C5', 'HQ'] },
  { p1: ['D2', 'D5'], p2: ['DJ', 'DA'], community: ['D3', 'D6', 'DT', 'C5', 'HQ'] },
  // Full House
  { p1: ['DQ', 'C2'], p2: ['CT', 'C4'], community: ['HQ', 'SQ', 'HT', 'DT', 'C3'] },
  { p1: ['ST', 'C2'], p2: ['CQ', 'C4'], community: ['HQ', 'SQ', 'HT', 'DT', 'C3'] },
  // Four of a Kind
  { p1: ['HA', 'S7'], p2: ['DJ', 'C5'], community: ['HT', 'ST', 'CT', 'DT', 'HK'] },
  { p1: ['C2', 'C3'], p2: ['C5', 'HK'], community: ['HT', 'ST', 'CT', 'DT', 'S8'] },
  // Straight Flush
  { p1: ['H7', 'HA'], p2: ['H2', 'SA'], community: ['H3', 'H4', 'H5', 'H6', 'HT'] },
  { p1: ['S6', 'C2'], p2: ['SJ', 'D5'], community: ['S7', 'S8', 'S9', 'ST', 'DK'] },
  // Royal Flush
  { p1: ['H2', 'C3'], p2: ['S4', 'C5'], community: ['DT', 'DJ', 'DQ', 'DK', 'DA'] },
];

function randomHand() {
  return HANDS[Math.floor(Math.random() * HANDS.length)];
}

function randomComparison() {
  return COMPARISONS[Math.floor(Math.random() * COMPARISONS.length)];
}

export default () => {
  client.connect('localhost:50051', { plaintext: true });

  // --- Test 1: EvaluateHand ---
  const hand = randomHand();
  const allCards = [...hand.hole, ...hand.community];

  const evalStart = Date.now();
  const evalResp = client.invoke('poker.PokerService/EvaluateHand', {
    cards: allCards,
  });
  evaluateDuration.add(Date.now() - evalStart);

  const evalOk = check(evalResp, {
    'evaluate: status OK': (r) => r && r.status === grpc.StatusOK,
    'evaluate: has hand rank': (r) => r && r.message && r.message.handRank !== '',
    'evaluate: has 5 best cards': (r) => r && r.message && r.message.bestFive.length === 5,
  });
  if (!evalOk) grpcErrors.add(1);

  // --- Test 2: CompareHands ---
  const cmp = randomComparison();

  const cmpStart = Date.now();
  const cmpResp = client.invoke('poker.PokerService/CompareHands', {
    player1_cards: [...cmp.p1, ...cmp.community],
    player2_cards: [...cmp.p2, ...cmp.community],
  });
  compareDuration.add(Date.now() - cmpStart);

  const cmpOk = check(cmpResp, {
    'compare: status OK': (r) => r && r.status === grpc.StatusOK,
    'compare: valid winner': (r) => r && r.message && [0, 1, 2].includes(r.message.winner),
    'compare: has description': (r) => r && r.message && r.message.description !== '',
  });
  if (!cmpOk) grpcErrors.add(1);

  // --- Test 3: CalculateWinProbability ---
  const probHand = randomHand();

  const probStart = Date.now();
  const probResp = client.invoke('poker.PokerService/CalculateWinProbability', {
    hole_cards: probHand.hole,
    community: [],         // Pre-flop: maximum simulation work
    num_opponents: 1,
    iterations: 1000,      // Keep iterations low for load testing
  });
  probabilityDuration.add(Date.now() - probStart);

  const probOk = check(probResp, {
    'probability: status OK': (r) => r && r.status === grpc.StatusOK,
    'probability: win >= 0': (r) => r && r.message && r.message.winProbability >= 0,
    'probability: sums to ~100': (r) => {
      if (!r || !r.message) return false;
      const total = r.message.winProbability + r.message.tieProbability + r.message.lossProbability;
      return total > 99 && total < 101;
    },
  });
  if (!probOk) grpcErrors.add(1);

  client.close();
  sleep(0.1);
};
