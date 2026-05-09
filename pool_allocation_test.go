package main

import (
	"math"
	"testing"

	"github.com/shopspring/decimal"
)

// InitializePool computes for each entity its Capacity, AvailableCapacity, and CurrentRatio.
func InitializePool(pool SubgraphsPool, totalTokensAllocated, totalTokensSignalled decimal.Decimal) {
	for i := range pool {
		// Protect against division by zero
		if totalTokensSignalled.IsZero() {
			pool[i].Capacity = decimal.Zero
			pool[i].AvailableCapacity = decimal.Zero
			pool[i].CurrentRatio = decimal.Zero
			pool[i].Amount = decimal.Zero
			continue
		}

		// Capacity = totalTokensAllocated * (SignalledTokens / totalTokensSignalled)
		pool[i].Capacity = totalTokensAllocated.Mul(pool[i].SignalledTokens.Div(totalTokensSignalled))

		// AvailableCapacity = Capacity - StakedTokens
		pool[i].AvailableCapacity = pool[i].Capacity.Sub(pool[i].StakedTokens)

		// CurrentRatio = StakedTokens / Capacity
		if pool[i].Capacity.IsZero() {
			pool[i].CurrentRatio = decimal.Zero
		} else {
			pool[i].CurrentRatio = pool[i].StakedTokens.Div(pool[i].Capacity)
		}

		// Initialize Amount to zero
		pool[i].Amount = decimal.Zero
	}
}

func TestAllocateTokensCases(t *testing.T) {
	type testCase struct {
		name                 string
		stakeAmount          decimal.Decimal
		totalTokensAllocated decimal.Decimal
		totalTokensSignalled decimal.Decimal
		pool                 SubgraphsPool
		// expectedAllocations maps SubgraphHash to expected final StakedTokens.
		expectedAllocations map[string]decimal.Decimal
		// expectedTokensLeft is the expected tokens left unused.
		expectedTokensLeft decimal.Decimal
	}

	tests := []testCase{
		{
			name:                 "Normal allocation with two subgraphs",
			stakeAmount:          decimal.NewFromInt(20),
			totalTokensAllocated: decimal.NewFromInt(100),
			totalTokensSignalled: decimal.NewFromInt(100),
			pool: SubgraphsPool{
				{SubgraphHash: "A", SignalledTokens: decimal.NewFromInt(40), StakedTokens: decimal.NewFromInt(20), Enabled: true},
				{SubgraphHash: "B", SignalledTokens: decimal.NewFromInt(40), StakedTokens: decimal.NewFromInt(20), Enabled: true},
			},
			// Expected outcome (simulation with our tie-breaking):
			// For example, final: A = 53, B = 37.
			expectedAllocations: map[string]decimal.Decimal{
				"A": decimal.NewFromInt(30),
				"B": decimal.NewFromInt(30),
			},
			expectedTokensLeft: decimal.Zero,
		},
		{
			name:                 "Stake amount exceeds available capacity (overfilling allowed)",
			stakeAmount:          decimal.NewFromInt(40),
			totalTokensAllocated: decimal.NewFromInt(100),
			totalTokensSignalled: decimal.NewFromInt(100),
			pool: SubgraphsPool{
				{SubgraphHash: "A", SignalledTokens: decimal.NewFromInt(60), StakedTokens: decimal.NewFromInt(40), Enabled: true},
				{SubgraphHash: "B", SignalledTokens: decimal.NewFromInt(60), StakedTokens: decimal.NewFromInt(40), Enabled: true},
			},
			// Expected outcome (simulation): A = 62, B = 38.
			expectedAllocations: map[string]decimal.Decimal{
				"A": decimal.NewFromInt(60),
				"B": decimal.NewFromInt(60),
			},
			expectedTokensLeft: decimal.Zero,
		},
		{
			name:                 "Subgraph overfilled initially",
			stakeAmount:          decimal.NewFromInt(40),
			totalTokensAllocated: decimal.NewFromInt(100),
			totalTokensSignalled: decimal.NewFromInt(100),
			pool: SubgraphsPool{
				{SubgraphHash: "A", SignalledTokens: decimal.NewFromInt(40), StakedTokens: decimal.NewFromInt(50), Enabled: true}, // Capacity=100, ratio=1.20
				{SubgraphHash: "B", SignalledTokens: decimal.NewFromInt(40), StakedTokens: decimal.NewFromInt(50), Enabled: true}, // Capacity=40, ratio=0.50
			},
			// Expected: A remains 120, B becomes 50.
			expectedAllocations: map[string]decimal.Decimal{
				"A": decimal.NewFromInt(70),
				"B": decimal.NewFromInt(70),
			},
			expectedTokensLeft: decimal.Zero,
		},
		{
			name:                 "Stake amount zero",
			stakeAmount:          decimal.Zero,
			totalTokensAllocated: decimal.NewFromInt(100),
			totalTokensSignalled: decimal.NewFromInt(100),
			pool: SubgraphsPool{
				{SubgraphHash: "A", SignalledTokens: decimal.NewFromInt(60), StakedTokens: decimal.NewFromInt(40), Enabled: true},
				{SubgraphHash: "B", SignalledTokens: decimal.NewFromInt(40), StakedTokens: decimal.NewFromInt(20), Enabled: true},
			},
			expectedAllocations: map[string]decimal.Decimal{
				"A": decimal.NewFromInt(40),
				"B": decimal.NewFromInt(20),
			},
			expectedTokensLeft: decimal.Zero,
		},
		{
			name:                 "Multiple subgraphs (3 subgraphs)",
			stakeAmount:          decimal.NewFromInt(30),
			totalTokensAllocated: decimal.NewFromInt(100),
			totalTokensSignalled: decimal.NewFromInt(100),
			// Subgraph A: Allocated 30, Signalled 50 → Capacity = 50, ratio = 0.60.
			// Subgraph B: Allocated 40, Signalled 30 → Capacity = 30, ratio ≈ 1.33 (overfilled).
			// Subgraph C: Allocated 10, Signalled 20 → Capacity = 20, ratio = 0.50.
			pool: SubgraphsPool{
				{SubgraphHash: "A", SignalledTokens: decimal.NewFromInt(30), StakedTokens: decimal.NewFromInt(20), Enabled: true},
				{SubgraphHash: "B", SignalledTokens: decimal.NewFromInt(30), StakedTokens: decimal.NewFromInt(40), Enabled: true},
				{SubgraphHash: "C", SignalledTokens: decimal.NewFromInt(30), StakedTokens: decimal.NewFromInt(10), Enabled: true},
			},
			// Expected: Only A and C receive tokens.
			// One simulation may yield: A = 39, B remains 40, C = 21.
			expectedAllocations: map[string]decimal.Decimal{
				"A": decimal.NewFromInt(30),
				"B": decimal.NewFromInt(40),
				"C": decimal.NewFromInt(30),
			},
			expectedTokensLeft: decimal.Zero,
		},
		{
			name:                 "Multiple subgraphs (3 subgraphs) last overfilled",
			stakeAmount:          decimal.NewFromInt(70),
			totalTokensAllocated: decimal.NewFromInt(100),
			totalTokensSignalled: decimal.NewFromInt(100),
			// Subgraph A: Allocated 30, Signalled 50 → Capacity = 50, ratio = 0.60.
			// Subgraph B: Allocated 40, Signalled 30 → Capacity = 30, ratio ≈ 1.33 (overfilled).
			// Subgraph C: Allocated 10, Signalled 20 → Capacity = 20, ratio = 0.50.
			pool: SubgraphsPool{
				{SubgraphHash: "A", SignalledTokens: decimal.NewFromInt(30), StakedTokens: decimal.NewFromInt(20), Enabled: true},
				{SubgraphHash: "B", SignalledTokens: decimal.NewFromInt(30), StakedTokens: decimal.NewFromInt(10), Enabled: true},
				{SubgraphHash: "C", SignalledTokens: decimal.NewFromInt(30), StakedTokens: decimal.NewFromInt(50), Enabled: true},
			},
			// Expected: Only A and C receive tokens.
			// One simulation may yield: A = 39, B remains 40, C = 21.
			expectedAllocations: map[string]decimal.Decimal{
				"A": decimal.NewFromInt(50),
				"B": decimal.NewFromInt(50),
				"C": decimal.NewFromInt(50),
			},
			expectedTokensLeft: decimal.Zero,
		},
		{
			name:                 "Zero stake amount with mixed capacity utilization",
			stakeAmount:          decimal.Zero,
			totalTokensAllocated: decimal.NewFromInt(100),
			totalTokensSignalled: decimal.NewFromInt(100),
			pool: SubgraphsPool{
				{SubgraphHash: "A", SignalledTokens: decimal.NewFromInt(30), StakedTokens: decimal.NewFromInt(40), Enabled: true}, // overfilled
				{SubgraphHash: "B", SignalledTokens: decimal.NewFromInt(40), StakedTokens: decimal.NewFromInt(20), Enabled: true}, // underfilled
				{SubgraphHash: "C", SignalledTokens: decimal.NewFromInt(30), StakedTokens: decimal.NewFromInt(30), Enabled: true}, // exactly filled
			},
			expectedAllocations: map[string]decimal.Decimal{
				"A": decimal.NewFromInt(40),
				"B": decimal.NewFromInt(20),
				"C": decimal.NewFromInt(30),
			},
			expectedTokensLeft: decimal.Zero,
		},
	}

	// Allow a tolerance of 0.5 tokens (since token-by-token allocation order may vary).
	tolerance := 0.5

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Copy the pool
			poolCopy := make(SubgraphsPool, len(tc.pool))
			copy(poolCopy, tc.pool)
			InitializePool(poolCopy, tc.totalTokensAllocated, tc.totalTokensSignalled)
			finalPool, tokensLeft := AllocateTokens(poolCopy, tc.stakeAmount)
			if math.Abs(tokensLeft.InexactFloat64()-tc.expectedTokensLeft.InexactFloat64()) > tolerance {
				t.Errorf("Expected tokens left %s, got %s", tc.expectedTokensLeft.StringFixed(2), tokensLeft.StringFixed(2))
			}
			for _, entity := range finalPool {
				exp, ok := tc.expectedAllocations[entity.SubgraphHash]
				if !ok {
					t.Errorf("Unexpected subgraph ID %s", entity.SubgraphHash)
					continue
				}
				diff := math.Abs(entity.StakedTokens.InexactFloat64() - exp.InexactFloat64())
				if diff > tolerance {
					t.Errorf("Subgraph %s expected allocated %s, got %s (diff %.2f)",
						entity.SubgraphHash, exp.StringFixed(2), entity.StakedTokens.StringFixed(2), diff)
				}
			}
		})
	}
}
