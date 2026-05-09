package main

import (
	"fmt"
	"sort"

	"github.com/shopspring/decimal"
)

// AllocateTokens allocates tokens one by one from stakeAmount to the pool entities.
// It prioritizes entities with the highest CurrentRatio, which represents subgraphs
// that have more signals relative to their stake (i.e., are under-allocated).
// In case of ties (within epsilon), it uses a round-robin order based on SubgraphHash (descending).
// Overfilling is allowed (i.e. CurrentRatio can go below 1), and the fields are updated accordingly.
// The function returns the updated pool and the number of tokens left unused.
func AllocateTokens(pool SubgraphsPool, stakeAmount decimal.Decimal) (SubgraphsPool, decimal.Decimal) {
	// extra tokens to allocate
	extra := stakeAmount
	one := decimal.NewFromInt(100)
	totalAllocated := decimal.Zero

	// Loop until extra is 0.
	for extra.GreaterThan(decimal.Zero) {
		// Sort pool by CurrentRatio descending; if equal (within epsilon), sort by SubgraphHash descending.
		sort.Slice(pool, func(i, j int) bool {
			// Skip subgraphs with no signaled tokens
			hasSignalI := pool[i].SignalledTokens.GreaterThan(decimal.Zero)
			hasSignalJ := pool[j].SignalledTokens.GreaterThan(decimal.Zero)

			// If one has signals and the other doesn't, prioritize the one with signals
			if hasSignalI != hasSignalJ {
				return hasSignalI
			}

			ratioI := pool[i].CurrentRatio
			ratioJ := pool[j].CurrentRatio
			if ratioI.Cmp(ratioJ) == 0 {
				// Descending order on SubgraphHash.
				return pool[i].SubgraphHash > pool[j].SubgraphHash
			}
			return ratioI.Cmp(ratioJ) > 0 // Changed to > for descending order (highest first)
		})

		// Find the first entity with signaled tokens
		firstValidIndex := -1
		for i, entity := range pool {
			if entity.SignalledTokens.GreaterThan(decimal.Zero) {
				firstValidIndex = i
				break
			}
		}

		// If no valid entities found, return
		if firstValidIndex == -1 {
			tokensLeft := stakeAmount.Sub(totalAllocated)
			return pool, tokensLeft
		}

		highestRatio := pool[firstValidIndex].CurrentRatio // Changed from lowestRatio

		// Gather indices with ratio equal to highestRatio (within a small epsilon) and with signaled tokens
		var highestIndices []int // Changed from lowestIndices
		for i, entity := range pool {
			if !entity.SignalledTokens.GreaterThan(decimal.Zero) {
				continue
			}
			diff := entity.CurrentRatio.Sub(highestRatio).Abs()
			if diff.Cmp(decimal.NewFromFloat(0.0001)) <= 0 {
				highestIndices = append(highestIndices, i)
			} else {
				break
			}
		}

		// Round-robin allocation among entities with highest ratio.
		for _, idx := range highestIndices {
			if extra.Cmp(decimal.Zero) <= 0 {
				break
			}
			// Allocate one token to pool[idx]
			pool[idx].StakedTokens = pool[idx].StakedTokens.Add(one)
			pool[idx].Amount = pool[idx].Amount.Add(one)
			totalAllocated = totalAllocated.Add(one)
			extra = extra.Sub(one)

			// Recalculate CurrentRatio and AvailableCapacity.
			// CurrentRatio should be SignalledTokens/StakedTokens ratio
			if pool[idx].StakedTokens.IsZero() {
				// Avoid division by zero
				pool[idx].CurrentRatio = decimal.NewFromInt(999) // Effectively infinite ratio
			} else {
				// Recalculate as SignaledRatio/StakedRatio
				pool[idx].CurrentRatio = pool[idx].SignalledTokens.Div(pool[idx].StakedTokens)
			}

			pool[idx].AvailableCapacity = pool[idx].Capacity.Sub(pool[idx].StakedTokens)
		}
	}

	tokensLeft := stakeAmount.Sub(totalAllocated)
	return pool, tokensLeft
}

// PrintPool prints the state of the pool.
func PrintPool(pool SubgraphsPool) {
	for _, entity := range pool {
		fmt.Printf("Subgraph %s: StakedTokens=%s, Capacity=%s, AvailableCapacity=%s, CurrentRatio=%s, Amount=%s, SignalledTokens=%s, Enabled=%t\n",
			entity.SubgraphHash,
			entity.StakedTokens.StringFixed(2),
			entity.Capacity.StringFixed(2),
			entity.AvailableCapacity.StringFixed(2),
			entity.CurrentRatio.StringFixed(2),
			entity.Amount.StringFixed(2),
			entity.SignalledTokens.StringFixed(2),
			entity.Enabled,
		)
	}
}
